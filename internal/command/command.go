package command

import (
	"github.com/urfave/cli"
)

type actionFunction func()
type actionFunctionArgs func(...string)
type Command struct {
	Name        string
	Description string
	Category    string
	UsageText   string
}

func (command Command) AddCommand(fn actionFunction) {
	newCommand := cli.Command{
		Name:      command.Name,
		Usage:     command.Description,
		Category:  command.Category,
		UsageText: command.UsageText,
		Action: func(c *cli.Context) error {
			fn()
			return nil
		},
	}

	commands = append(commands, newCommand)
}

func (command Command) AddCommandWithArgs(fn actionFunctionArgs) {
	newCommand := cli.Command{
		Name:      command.Name,
		Usage:     command.Description,
		Category:  command.Category,
		UsageText: command.UsageText,
		Action: func(c *cli.Context) error {
			fn(c.Args()...)
			return nil
		},
	}

	commands = append(commands, newCommand)
}
