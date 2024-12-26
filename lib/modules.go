package lib

import "context"

type Module interface {
	Tag() string
	Status() string
	Init(ctx context.Context) error
	MustInit(ctx context.Context)
	Stop(ctx context.Context) error
	MustStop(ctx context.Context)
}
