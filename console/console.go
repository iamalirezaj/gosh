package console

import "gosh/console/commands"
import "github.com/codegangsta/cli"
import "fmt"
import "reflect"

type Console struct {
	Name string
	Version string
	Commands []commands.CommandInterface
}

func (console Console) AddCommand(command commands.CommandInterface) Console {

	console.Commands = append(console.Commands, command)
	return console
}

func (console Console) AddCommands(commands []commands.CommandInterface) Console {

	for _, command := range commands {
		console = console.AddCommand(command)
	}

	return console
}

func (console Console) Run(args cli.Args) {

	application := cli.NewApp()
	application.Name = console.Name
	application.Version = console.Version
	application.Commands = []cli.Command{}
	application.Flags = append(application.Flags, []cli.Flag{}...)

	for _, command := range console.Commands {

		cmd := command.Configure()

		cmd.Action = func(context *cli.Context) error {

			for _, flag := range cmd.Flags {
				if reflect.TypeOf(flag).Name() == "RequiredFlag" {
					if context.Args().Present() {
						command.Handle(context)
					} else {
						return cli.NewExitError(
							fmt.Sprintf(
								"Missing %v argument for '%v'",
								flag.GetName(), context.Command.Name,
							), 3,
						)
					}
				} else {
					continue
				}
			}

			return nil
		}
		application.Commands = append(application.Commands, cmd)
	}

	application.Run(args)
}