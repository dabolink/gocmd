package command

import (
	"context"
)

type CommandRunnable[TProvider any] interface {
	Run(ctx context.Context, in TProvider) error
}

type CommandData[TInput CommandInput, TProvider any] interface {
	GetInfo() CommandInfo
	Get(TInput) (CommandRunnable[TProvider], error)
}

type CommandInput interface {
	GetCommandName() string
}

type CommandInfo struct {
	IsUserVisible bool
	Name          string
	Aliases       []string
}

type CommandDataRetriever interface {
	GetInfo() CommandInfo
}
