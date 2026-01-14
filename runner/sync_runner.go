package runner

import (
	"context"
)

type SyncRunner[TProvider any] struct {
	runner func(context.Context, Runnable[TProvider]) error
	errCh  chan error
}

func (r *SyncRunner[TProvider]) Run(ctx context.Context, runnable Runnable[TProvider]) error {
	return r.runner(ctx, runnable)
}

func (r *SyncRunner[_]) Init() {}

func (r *SyncRunner[_]) Wait() {}

func NewSyncRunner[TProvider any](runnerFn func(context.Context, Runnable[TProvider]) error) *SyncRunner[TProvider] {
	return &SyncRunner[TProvider]{
		runner: runnerFn,
		errCh:  make(chan error, 1),
	}
}
