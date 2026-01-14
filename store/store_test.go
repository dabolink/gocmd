package store

import (
	"reflect"
	"testing"

	"github.com/dabolink/gocmd/command"
)

type SimpleCommandInput struct {
	cmdName string
}

func (in *SimpleCommandInput) GetCommandName() string {
	return in.cmdName
}

type SimpleCommandData struct {
	command.CommandInfo
}

func (d *SimpleCommandData) Get(*SimpleCommandInput) (command.CommandRunnable[*SimpleCommandData], error) {
	return nil, nil
}

func (d *SimpleCommandData) GetInfo() command.CommandInfo {
	return d.CommandInfo
}

func TestCommandStore_get(t *testing.T) {
	type fields struct {
		commands map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]
	}
	type args struct {
		cmdName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    command.CommandData[*SimpleCommandInput, *SimpleCommandData]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommandStore[*SimpleCommandInput, *SimpleCommandData]{
				commands: tt.fields.commands,
			}
			got, err := s.get(tt.args.cmdName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandStore_Add(t *testing.T) {
	type fields struct {
		commands map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]
	}
	type args struct {
		data command.CommandData[*SimpleCommandInput, *SimpleCommandData]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommandStore[*SimpleCommandInput, *SimpleCommandData]{
				commands: tt.fields.commands,
			}
			s.Add(tt.args.data)
		})
	}
}

func TestCommandStore_List(t *testing.T) {
	type fields struct {
		commands map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "no commands",
			fields: fields{
				commands: map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]{},
			},
			want: []string{},
		},
		{
			name: "has visible commands",
			fields: fields{
				commands: map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]{
					"key": &SimpleCommandData{
						CommandInfo: command.CommandInfo{
							Name:          "test",
							IsUserVisible: true,
							Aliases:       []string{},
						},
					},
				},
			},
			want: []string{"test"},
		},
		{
			name: "has invisble commands",
			fields: fields{
				commands: map[string]command.CommandData[*SimpleCommandInput, *SimpleCommandData]{
					"key": &SimpleCommandData{
						command.CommandInfo{
							Name:          "test",
							IsUserVisible: false,
							Aliases:       []string{},
						},
					},
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommandStore[*SimpleCommandInput, *SimpleCommandData]{
				commands: tt.fields.commands,
			}
			if got := s.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
