package command

import (
	"os"

	"github.com/urfave/cli"
	"fmt"
)

var commands []cli.Command
var app *cli.App

func init() {
	app = cli.NewApp()
}

func RunApp() {
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("App run failed, error: ", err)
	}
}
