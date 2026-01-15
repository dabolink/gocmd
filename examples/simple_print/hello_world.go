package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dabolink/gocmd"
	"github.com/dabolink/gocmd/command"
	"github.com/dabolink/gocmd/parser"
	"github.com/dabolink/gocmd/runner"
	"github.com/dabolink/gocmd/store"
)

type SimpleCommandInput struct {
	CommandName string
}

func (in *SimpleCommandInput) GetCommandName() string {
	return in.CommandName
}

func main() {
	cmd := gocmd.MakeCommand(
		command.CommandInfo{
			IsUserVisible: true,
			Name:          "print",
			Aliases:       []string{"helloworld"},
		},
		parser.Identity,
		func(ctx context.Context, globals any, in *SimpleCommandInput) error {
			fmt.Printf("printing Hello, World! from %s\n", in.GetCommandName())
			return nil
		})

	cli := gocmd.NewClient(
		store.NewConfig(cmd),
		runner.Config[any]{
			Provider:      struct{}{},
			RunInParallel: false,
		},
	)

	var cmdName string
	if len(os.Args) == 2 {
		cmdName = os.Args[1]
	} else {
		fmt.Println("Usage: go run <file> <command_name>")
	}
	err := cli.Run(context.TODO(), &SimpleCommandInput{CommandName: cmdName})
	if err != nil {
		if err == store.ErrUnknownCommand {
			fmt.Printf("unknown command: %s\n", cmdName)
			return
		}
		panic(err)
	}
	cli.Wait()
}
