package gocmd

import (
	"context"

	"github.com/dabolink/gocmd/command"
	"github.com/dabolink/gocmd/parser"
	"github.com/dabolink/gocmd/runner"
)

type Command[TInput command.CommandInput, T any, TProvider any] struct {
	parser  parser.ParserFn[TInput, T]
	fn      func(T) runner.Runnable[TProvider]
	cmdInfo command.CommandInfo
}

func (c *Command[TInput, _, TProvider]) Get(args TInput) (command.CommandRunnable[TProvider], error) {
	in, err := c.parser(args)
	if err != nil {
		return nil, err
	}
	return c.fn(in), nil
}

func (c *Command[_, _, _]) GetInfo() command.CommandInfo {
	return c.cmdInfo
}

type runnable[TInput any] struct {
	runFn func(ctx context.Context, input TInput) error
}

func (r *runnable[TProvider]) Run(ctx context.Context, provider TProvider) error {
	return r.runFn(ctx, provider)
}
