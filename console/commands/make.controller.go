package commands

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"os"
	"reflect"
	"io/ioutil"
	"strings"
	"fmt"
)

type MakeControllerCommand Command

func (c MakeControllerCommand) Configure(command cli.Command) cli.Command {

	command.Name        =   "make:controller"
	command.UsageText   =   "gosh make:controller {controller_name} "
	command.Description =   `Generate a new controller`
	command.Flags       =   []cli.Flag {
		cli.StringFlag {
			Name: "name",
		},
	}
	command.Action      =   func(context *cli.Context) error {

		if context.Args().First() == "" {

			return cli.NewExitError(
				fmt.Sprintf(
					"Missing %v argument for '%v'",
					"name", context.Command.Name,
				), 3,
			)
		}

		currentPath := os.Getenv("PWD")

		controller := context.Args().First()

		file := currentPath + "/" + controller + ".controller.go"

		// detect if file exists
		var _, controllerFile = os.Stat(file)

		// create file if not exists
		if os.IsNotExist(controllerFile) {

			stubfile := os.Getenv("GOPATH") + "/src/" + reflect.TypeOf(c).PkgPath() + "/stubs/controller.stub"
			stubContent, err := ioutil.ReadFile(stubfile)

			if err == nil {
				contents := strings.Replace(string(stubContent), "{CONTROLLER}", controller, -1)
				err := ioutil.WriteFile(file, []byte(contents), 0644)

				if err == nil {
					color.Green("Controller created successfully.")
				}
			}
		} else {

			color.Red(controller + " Controller is exist")
		}

		return nil
	}

	return command
}

