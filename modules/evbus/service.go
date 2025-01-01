package evbus

import (
	"context"
	"errors"
	"reflect"
	"sync"

	"github.com/kijudev/blueprint/lib"
)

type Event interface{}

type Service struct {
	maxGoroutines uint

	registry map[reflect.Type]struct{}
	handlers map[reflect.Type][]reflect.Value

	wg   sync.WaitGroup
	lock sync.Mutex
}

func NewEventBusService(maxGoroutines uint) *Service {
	return &Service{
		maxGoroutines: maxGoroutines,
		registry:      make(map[reflect.Type]struct{}),
		handlers:      make(map[reflect.Type][]reflect.Value),
	}
}

func (s *Service) Register(ctx context.Context, ev Event) error {
	e := errors.New("evbus.Module.Service.Register")

	s.lock.Lock()
	defer s.lock.Unlock()

	evType := reflect.TypeOf(ev)

	if evType.Kind() == reflect.Ptr {
		evType = evType.Elem()
	}

	if _, ok := s.registry[evType]; ok {
		return lib.JoinErrors(e, errors.New("Event already registered"))
	}

	s.registry[evType] = struct{}{}

	return nil
}

func (s *Service) MustRegister(ctx context.Context, ev Event) {
	if err := s.Register(ctx, ev); err != nil {
		panic(err)
	}
}

func (s *Service) IsRegistered(ctx context.Context, ev Event) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	evType := reflect.TypeOf(ev)

	if evType.Kind() == reflect.Ptr {
		evType = evType.Elem()
	}

	if _, ok := s.registry[evType]; ok {
		return true
	}

	return false
}

func (s *Service) Subscribe(ctx context.Context, handler any) error {
	e := errors.New("evbus.Module.Service.Subscribe")

	s.lock.Lock()
	defer s.lock.Unlock()

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)

	if handlerType.Kind() != reflect.Func {
		return lib.JoinErrors(e, errors.New("Handler is not a function"))
	}

	numArgs := handlerType.NumIn()

	if numArgs != 2 {
		return lib.JoinErrors(e, errors.New("Handler must have two arguments: ctx, event"))
	}

	handlerCtxType := handlerType.In(0)
	handlerEvType := handlerType.In(1)

	if !handlerCtxType.Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		return lib.JoinErrors(e, errors.New("The first handler argument in not context.Context"))
	}

	if handlerEvType.Kind() == reflect.Ptr {
		return lib.JoinErrors(e, errors.New("Handle must not accept events passed by pointers"))
	}

	if _, ok := s.registry[handlerEvType]; !ok {
		return lib.JoinErrors(e, errors.New("Event not registered"))
	}

	s.handlers[handlerEvType] = append(s.handlers[handlerEvType], handlerValue)

	return nil
}

func (s *Service) MustSubscribe(ctx context.Context, handler any) {
	if err := s.Subscribe(ctx, handler); err != nil {
		panic(err)
	}
}

func (s *Service) Dispatch(ctx context.Context, ev Event) error {
	e := errors.New("evbus.Module.Service.Dispatch")

	s.lock.Lock()
	defer s.lock.Unlock()

	evType := reflect.TypeOf(ev)
	evValue := reflect.ValueOf(ev)

	if evType.Kind() == reflect.Ptr && evValue.IsNil() {
		return lib.JoinErrors(e, errors.New("Event cannot be null"))
	}

	if evType.Kind() == reflect.Ptr {
		evType = evType.Elem()
		evValue = evValue.Elem()
	}

	if _, ok := s.registry[evType]; !ok {
		return lib.JoinErrors(e, errors.New("Event not registered"))
	}

	// handlers are insured to be valid
	handlers := s.handlers[evType]
	guard := make(chan struct{}, s.maxGoroutines)
	ctxValue := reflect.ValueOf(ctx)

	for _, handler := range handlers {
		s.wg.Add(1)
		guard <- struct{}{}

		go func(handler reflect.Value) {
			handler.Call([]reflect.Value{
				ctxValue,
				evValue,
			})

			s.wg.Done()
			<-guard
		}(handler)
	}

	return nil
}

func (s *Service) MustDispatch(ctx context.Context, ev Event) {
	if err := s.Dispatch(ctx, ev); err != nil {
		panic(err)
	}
}

func (s *Service) DispatchSync(ctx context.Context, ev Event) error {
	s.Wait(ctx)

	e := errors.New("evbus.Module.Service.DispatchSync")
	err := s.Dispatch(ctx, ev)

	if err != nil {
		return lib.JoinErrors(e, err)
	}

	s.Wait(ctx)
	return nil
}

func (s *Service) MustDispatchSync(ctx context.Context, ev Event) {
	if err := s.DispatchSync(ctx, ev); err != nil {
		panic(err)
	}
}

func (s *Service) Wait(ctx context.Context) {
	s.wg.Wait()
}
