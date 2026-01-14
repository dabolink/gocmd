package runner

import (
	"context"
	"fmt"
	"sync"
)

type Config[TProvider any] struct {
	Provider      TProvider
	RunInParallel bool
}

type Runner[TProvider any] struct {
	runner        func(context.Context, Runnable[TProvider]) error
	runInParallel bool
	wg            *sync.WaitGroup
	errCh         chan error
	closeOnce     sync.Once
}

func (r *Runner[_]) Init() {
	if !r.runInParallel {
		return
	}
	go func() {
		for err := range r.errCh {
			if err != nil {
				fmt.Printf("\n%v\n> ", err)
			}
		}
	}()
}

func (r *Runner[TProvider]) Run(ctx context.Context, runnable Runnable[TProvider]) error {
	r.run(ctx, runnable)
	if r.runInParallel {
		return nil
	}
	return <-r.errCh
}

func (r *Runner[_]) Wait() {
	r.wg.Wait()
	r.closeOnce.Do(func() {
		close(r.errCh)
	})
}

func (r *Runner[_]) Close() {
	r.wg.Wait()
	r.closeOnce.Do(func() { close(r.errCh) })
}
func (r *Runner[TProvider]) run(ctx context.Context, runnable Runnable[TProvider]) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		r.errCh <- r.runner(ctx, runnable)
	}()
}

func WithConfig[TProvider any](config Config[TProvider]) *Runner[TProvider] {
	return &Runner[TProvider]{
		runner: func(ctx context.Context, runnable Runnable[TProvider]) error {
			return runnable.Run(ctx, config.Provider)
		},
		wg:            &sync.WaitGroup{},
		runInParallel: config.RunInParallel,
		errCh:         make(chan error, 1),
	}
}
