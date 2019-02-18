package console

import (
	"github.com/codegangsta/cli"
	"github.com/goshco/gosh/console/commands"
)

type Console struct {
	Name string
	Version string
	Commands []commands.CommandInterface
}

func (console Console) AddCommands(commands []commands.CommandInterface) Console {

	console.Commands = append(console.Commands, commands...)
	return console
}

func (console Console) Run(args cli.Args) {

	application := cli.NewApp()
	application.Name = "Gosh"
	application.Version = "0.0.1"

	for _, command := range console.Commands {
		application.Commands = append(application.Commands, command.Configure(cli.Command{}))
	}

	application.Run(args)
}