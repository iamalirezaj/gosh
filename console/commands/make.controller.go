package commands

import (
	"github.com/codegangsta/cli"
	"gosh/console/commands/flags"
	"github.com/fatih/color"
	"os"
	"reflect"
	"io/ioutil"
	"strings"
)

type HelpCommand Command

func (command HelpCommand) Configure() cli.Command {

	return cli.Command {
		Name:        "make:controller",
		UsageText:   "gosh make:controller {controller_name} ",
		Description: `Generate a new controller`,
		Flags: []cli.Flag {
			flags.RequiredFlag{
				cli.StringFlag{
					Name: "name",
				},
			},
		},
	}
}

func (command HelpCommand) Handle(context *cli.Context) {

	currentPath := os.Getenv("PWD")

	controller := context.Args().First()

	file := currentPath + "/" + controller + ".controller.go"

	// detect if file exists
	var _, controllerFile = os.Stat(file)

	// create file if not exists
	if os.IsNotExist(controllerFile) {

		command.WriteFile(controller, file)
	} else {

		color.Red(controller + " Controller is exist")
	}
}

func (command HelpCommand) WriteFile(controller string,filepath string)  {

	stubfile := os.Getenv("GOPATH") + "/src/" + reflect.TypeOf(command).PkgPath() + "/stubs/controller.stub"

	content, err := ioutil.ReadFile(stubfile)
	if err == nil {

		newContents := strings.Replace(string(content), "{CONTROLLER}", controller, -1)

		err := ioutil.WriteFile(filepath, []byte(newContents), 0644)
		if err == nil {
			color.Green("Controller created successfully.")
		}
	}
}

