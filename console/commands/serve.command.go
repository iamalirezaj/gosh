package commands

import (
	"os/exec"
	"github.com/fatih/color"
	"github.com/codegangsta/cli"
)

type ServeCommand Command

func (c ServeCommand) Configure(command cli.Command) cli.Command {

	command.Name        =   "serve"
	command.Aliases     =   []string{"s"}
	command.UsageText   =   "gosh make:controller {controller_name} "
	command.Description =   `Generate a new controller`
	command.Flags 		=	[]cli.Flag {
		cli.StringFlag{
			Name: "host",
		},
		cli.StringFlag{
			Name: "port",
		},
	}

	command.Action      =   func(context *cli.Context) {

		Host := c.GetHost(context)
		Port := c.GetPort(context)

		cmd := exec.Command("go","run", "main.go", "welcome.controller.go", "router.provider.go", Host, Port)
		color.Green("Gosh development server started: <http://" + Host + ":" + Port + ">")
		output, _ := cmd.Output()
		color.Red(string(output))
	}

	return command
}

func (c ServeCommand) GetHost(context *cli.Context) string {

	if context.Args().First() == "" {
		return "127.0.0.1"
	}

	return context.Args().First()
}

func (c ServeCommand) GetPort(context *cli.Context) string {

	if context.Args().Get(1) == "" {
		return "8080"
	}

	return context.Args().Get(1)
}