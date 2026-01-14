package runner

import (
	"context"
	"fmt"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type MinimalContext struct {
	context.Context
	Logger
}

type logger struct {
}

func (l logger) Info(msg string) {
	fmt.Println("[info] " + msg)
}

func (l logger) Error(msg string) {
	fmt.Println("[error] " + msg)
}

func NewContext(ctx context.Context) MinimalContext {
	return MinimalContext{
		Logger:  logger{},
		Context: ctx,
	}
}
