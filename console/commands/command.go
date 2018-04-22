package commands

import "github.com/codegangsta/cli"

type Command struct {
	CommandInterface
}

type CommandInterface interface {
	Configure(command cli.Command) cli.Command
}