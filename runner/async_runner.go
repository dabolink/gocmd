package runner

import (
	"context"
	"fmt"
	"sync"
)

type AsyncRunner[TProvider any] struct {
	runner    func(context.Context, Runnable[TProvider]) error
	wg        *sync.WaitGroup
	errCh     chan error
	closeOnce sync.Once
}

func (r *AsyncRunner[_]) Init() {
	go func() {
		for err := range r.errCh {
			if err != nil {
				fmt.Printf("\n%v\n> ", err)
			}
		}
	}()
}

func (r *AsyncRunner[TProvider]) Run(ctx context.Context, runnable Runnable[TProvider]) error {
	r.wg.Go(func() {
		r.errCh <- r.runner(ctx, runnable)
	})
	return nil
}

func (r *AsyncRunner[_]) Wait() {
	r.wg.Wait()
	r.closeOnce.Do(func() {
		close(r.errCh)
	})
}

func NewAsyncRunner[TProvider any](runnerFn func(context.Context, Runnable[TProvider]) error) *AsyncRunner[TProvider] {
	return &AsyncRunner[TProvider]{
		runner: runnerFn,
		wg:     &sync.WaitGroup{},
		errCh:  make(chan error, 1),
	}
}
