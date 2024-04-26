package go_at_command

import (
	"fmt"
	"strconv"
	"strings"
)

type ATCommands struct {
	Commands []ATCommand
}

func (cmds ATCommands) String() string {
	var parts []string
	for _, cmd := range cmds.Commands {
		parts = append(parts, cmd.String())
	}
	return "AT" + strings.Join(parts, ";")
}

type ATCommand interface {
	fmt.Stringer

	GetCommandName() string
}

type ATTestCommand struct {
	CommandName string
}

func (cmd ATTestCommand) GetCommandName() string {
	return cmd.CommandName
}

func (cmd ATTestCommand) String() string {
	return cmd.CommandName + "=?"
}

type ATReadCommand struct {
	CommandName string
}

func (cmd ATReadCommand) GetCommandName() string {
	return cmd.CommandName
}

func (cmd ATReadCommand) String() string {
	return cmd.CommandName + "?"
}

type ATExecuteCommand struct {
	CommandName string
}

func (cmd ATExecuteCommand) String() string {
	return cmd.CommandName
}

func (cmd ATExecuteCommand) GetCommandName() string {
	return cmd.CommandName
}

type ATSetCommand struct {
	CommandName string
	Values      []string
}

func (cmd ATSetCommand) GetCommandName() string {
	return cmd.CommandName
}

func (cmd ATSetCommand) String() string {
	var values []string
	for _, value := range cmd.Values {
		if strings.ContainsAny(value, `", `) {
			values = append(values, strconv.Quote(value))
		} else {
			values = append(values, value)
		}
	}
	return cmd.CommandName + "=" + strings.Join(values, ",")
}
