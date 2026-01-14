package gocmd

import (
	"context"

	"github.com/dabolink/gocmd/command"
	"github.com/dabolink/gocmd/runner"
	"github.com/dabolink/gocmd/store"
)

var _ Interface[command.CommandInput, any] = &Client[command.CommandInput, any]{}

type Interface[TInput command.CommandInput, TProvider any] interface {
	Run(context.Context, TInput) error
	AddCommand(command.CommandData[TInput, TProvider])
	Wait()
}

type Client[TInput command.CommandInput, TProvider any] struct {
	commandStore *store.CommandStore[TInput, TProvider]
	runner       *runner.Runner[TProvider]
}

func (cli *Client[TInput, TProvider]) Run(ctx context.Context, in TInput) error {
	cmd, err := cli.commandStore.GetCommand(in)
	if err != nil {
		return err
	}
	return cli.runner.Run(ctx, cmd)
}

func (cli *Client[TInput, TProvider]) AddCommand(data command.CommandData[TInput, TProvider]) {
	cli.commandStore.Add(data)
}

func (cli *Client[_, _]) Wait() {
	cli.runner.Wait()
}

func MakeCommand[TInput command.CommandInput, T any, TProvider any](info command.CommandInfo, parser func(TInput) (T, error), runFn func(context.Context, TProvider, T) error) *Command[TInput, T, TProvider] {
	return &Command[TInput, T, TProvider]{
		parser: parser,
		fn: func(t T) runner.Runnable[TProvider] {
			return &runnable[TProvider]{
				runFn: func(ctx context.Context, in TProvider) error {
					return runFn(ctx, in, t)
				},
			}
		},
		cmdInfo: info,
	}
}

func NewClient[TInput command.CommandInput, TProvider any](storeConfig store.Config[TInput, TProvider], runnerConfig runner.Config[TProvider]) *Client[TInput, TProvider] {
	s := store.WithConfig(storeConfig)
	r := runner.WithConfig(runnerConfig)

	return &Client[TInput, TProvider]{
		commandStore: s,
		runner:       r,
	}
}
