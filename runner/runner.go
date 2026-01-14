package runner

import (
	"context"
)

type Runner[TProvider any] interface {
	Init()
	Run(context.Context, Runnable[TProvider]) error
	Wait()
}

type Config[TProvider any] struct {
	Provider      TProvider
	RunInParallel bool
}

func WithConfig[TProvider any](config Config[TProvider]) Runner[TProvider] {
	runnerFn := func(ctx context.Context, runnable Runnable[TProvider]) error {
		return runnable.Run(ctx, config.Provider)
	}

	if config.RunInParallel {
		return NewAsyncRunner(runnerFn)
	} else {
		return NewSyncRunner(runnerFn)
	}
}
