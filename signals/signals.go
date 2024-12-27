package signals

import (
	"context"
	"time"
)

type Event[T any] struct {
	Timestamp time.Time
	Data      T
}

type Subscriber[T any] func(context.Context, Event[T])

type Listen[T any] func(Subscriber[T])
type Dispatch[T any] func(data T)

type Signal[T any] struct {
	Listen   Listen[T]
	Dispatch Dispatch[T]

	subscribers []Subscriber[T]
}

func New[T any]() *Signal[T] {
	s := &Signal[T]{
		subscribers: make([]Subscriber[T], 0),
	}

	s.Listen = func(subscriber Subscriber[T]) {
		s.subscribers = append(s.subscribers, subscriber)
	}

	s.Dispatch = func(data T) {
		for _, subscriber := range s.subscribers {
			subscriber(context.Background(), Event[T]{
				Timestamp: time.Now().UTC(),
				Data:      data,
			})
		}
	}

	return s
}
