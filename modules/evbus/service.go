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

	registry map[string]reflect.Type
	handlers map[string][]reflect.Value

	wg   sync.WaitGroup
	lock sync.Mutex
}

func NewEventBusService(maxGoroutines uint) *Service {
	return &Service{
		maxGoroutines: maxGoroutines,
		registry:      make(map[string]reflect.Type),
		handlers:      make(map[string][]reflect.Value),
	}
}

func (s *Service) Register(ctx context.Context, evCode string, ev Event) error {
	e := errors.New("evbus.Module.Service.Register")

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.registry[evCode]; ok {
		return lib.JoinErrors(e, errors.New("Event code already registered"))
	}

	evType := reflect.TypeOf(ev)

	if evType.Kind() == reflect.Ptr {
		evType = evType.Elem()
	}

	s.registry[evCode] = evType

	return nil
}

func (s *Service) MustRegister(ctx context.Context, evCode string, ev Event) {
	if err := s.Register(ctx, evCode, ev); err != nil {
		panic(err)
	}
}

func (s *Service) IsRegistered(ctx context.Context, evCode string) bool {
	if _, ok := s.registry[evCode]; ok {
		return true
	}

	return false
}

func (s *Service) Subscribe(ctx context.Context, evCode string, handler any) error {
	e := errors.New("evbus.Module.Service.Subscribe")

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.registry[evCode]; !ok {
		return lib.JoinErrors(e, errors.New("Event code not registered"))
	}

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)

	if handlerType.Kind() != reflect.Func {
		return lib.JoinErrors(e, errors.New("Handler is not a function"))
	}

	numArgs := handlerType.NumIn()
	if numArgs != 2 {
		return lib.JoinErrors(e, errors.New("Handler must have two arguments: ctx, event"))
	}

	if !handlerType.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		return lib.JoinErrors(e, errors.New("The first handler argument in not context.Context"))
	}

	if handlerType.In(1).Kind() == reflect.Ptr {
		return lib.JoinErrors(e, errors.New("Handle must not accept events passed by pointers"))
	}

	if handlerType.In(1) != s.registry[evCode] {
		return lib.JoinErrors(e, errors.New("Event type mismatch"))
	}

	s.handlers[evCode] = append(s.handlers[evCode], handlerValue)

	return nil
}

func (s *Service) MustSubscribe(ctx context.Context, evCode string, handler any) {
	if err := s.Subscribe(ctx, evCode, handler); err != nil {
		panic(err)
	}
}

func (s *Service) Dispatch(ctx context.Context, evCode string, ev Event) error {
	e := errors.New("evbus.Module.Service.Dispatch")

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.registry[evCode]; !ok {
		return lib.JoinErrors(e, errors.New("Event code not registered"))
	}

	evType := reflect.TypeOf(ev)
	evValue := reflect.ValueOf(ev)

	if evType.Kind() == reflect.Pointer && evValue.IsNil() {
		return lib.JoinErrors(e, errors.New("Event cannot be null"))
	}

	if evType.Kind() == reflect.Pointer {
		evType = evType.Elem()
		evValue = evValue.Elem()
	}

	if t, _ := s.registry[evCode]; evType != t {
		return lib.JoinErrors(e, errors.New("Event type mismatch"))
	}

	// handlers are insured to be valid
	handlers := s.handlers[evCode]
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

func (s *Service) MustDispatch(ctx context.Context, evCode string, ev any) {
	if err := s.Dispatch(ctx, evCode, ev); err != nil {
		panic(err)
	}
}

func (s *Service) Wait(ctx context.Context) {
	s.wg.Wait()
}
