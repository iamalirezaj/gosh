package commands

import "github.com/codegangsta/cli"

type Command struct {
	CommandInterface
}

type CommandInterface interface {
	Configure() cli.Command
	Handle(c *cli.Context)
}