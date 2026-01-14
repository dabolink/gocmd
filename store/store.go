package store

import (
	"errors"

	"github.com/dabolink/gocmd/command"
)

var (
	ErrUnknownCommand = errors.New("unknown command")
)

type Config[TInput command.CommandInput, TProvider any] struct {
	Commands []command.CommandData[TInput, TProvider]
}

func NewConfig[TInput command.CommandInput, TProvider any](commands ...command.CommandData[TInput, TProvider]) Config[TInput, TProvider] {
	return Config[TInput, TProvider]{
		Commands: commands,
	}
}

func Default[TInput command.CommandInput, TProvider any]() *CommandStore[TInput, TProvider] {
	return &CommandStore[TInput, TProvider]{
		commands: make(map[string]command.CommandData[TInput, TProvider], 10),
		aliases:  make(map[string]string, 10),
	}
}

func WithConfig[TInput command.CommandInput, TProvider any](config Config[TInput, TProvider]) *CommandStore[TInput, TProvider] {
	s := &CommandStore[TInput, TProvider]{
		commands: make(map[string]command.CommandData[TInput, TProvider], len(config.Commands)),
		aliases:  make(map[string]string, len(config.Commands)),
	}
	s.Add(config.Commands...)
	return s
}

type CommandStore[TInput command.CommandInput, TProvider any] struct {
	commands map[string]command.CommandData[TInput, TProvider]
	aliases  map[string]string
}

func (c *CommandStore[TInput, TProvider]) ListCommands() []command.CommandInfo {
	commands := make([]command.CommandInfo, 0, len(c.commands))
	for _, val := range c.commands {
		commands = append(commands, val.GetInfo())
	}
	return commands
}

func (s *CommandStore[_, _]) GetInfo(cmd string) (command.CommandInfo, error) {
	info, err := s.GetData(cmd)
	if err != nil {
		return command.CommandInfo{}, err
	}
	return info.GetInfo(), err
}

func (s *CommandStore[TInput, TProvider]) GetCommand(input TInput) (command.CommandRunnable[TProvider], error) {
	wrapper, err := s.get(input.GetCommandName())
	if err != nil {
		return nil, err
	}
	return wrapper.Get(input)
}

func (s *CommandStore[_, _]) GetData(cmdName string) (command.CommandDataRetriever, error) {
	return s.get(cmdName)
}

func (s *CommandStore[TInput, TProvider]) get(cmdName string) (command.CommandData[TInput, TProvider], error) {
	if result, ok := s.commands[cmdName]; ok {
		return result, nil
	} else if result, ok := s.aliases[cmdName]; ok {
		return s.get(result)
	}
	return nil, ErrUnknownCommand
}

func (s *CommandStore[TInput, TProvider]) Add(wrappers ...command.CommandData[TInput, TProvider]) {
	for _, wrapper := range wrappers {
		info := wrapper.GetInfo()
		s.commands[info.Name] = wrapper
		s.AddAliases(info.Name, info.Aliases...)
	}

}

func (s *CommandStore[_, _]) AddAliases(cmdName string, aliases ...string) {
	for _, alias := range aliases {
		s.aliases[alias] = cmdName
	}
}
