package evbus

import (
	"context"
	"sync"
)

type Service struct {
	maxGoroutines int

	handlers map[string][]func(context.Context, Event)

	wg   sync.WaitGroup
	lock sync.Mutex
}

type Event interface {
	Tag() string
}

func NewEventBusService(maxGoroutines int) *Service {
	return &Service{
		maxGoroutines: maxGoroutines,
	}
}

func (s *Service) Subscribe(evCode string, handler func(ctx context.Context, args ...any)) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.handlers[evCode]; !ok {
		s.handlers[evCode] = make([]func(context.Context, Event), 0)
	}

}

func (s *Service) Wait() {
	s.wg.Wait()
}
