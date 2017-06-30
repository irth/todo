package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var commandUpdate = cli.Command{
	Name:      "update",
	Aliases:   []string{"u"},
	Usage:     "update a TODO item",
	UsageText: "todo update -i ID [--category,-c CATEGORY] [todo description...]",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "category, c",
		},
		cli.StringFlag{
			Name:  "id, i",
			Usage: "ID of the TODO that you want to update",
		},
	},
	Category: "TODOs",

	Action: func(c *cli.Context) error {
		arg := c.String("id")
		id, err := strconv.Atoi(arg)
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID", 4)
		}

		todo, err := db.getTodo(id)
		if err != nil {
			return cli.NewExitError("Couldn't find the TODO", 4)
		}

		if c.IsSet("category") {
			todo.Category = c.String("category")
		}

		text := strings.TrimSpace(strings.Join(c.Args(), " "))
		if len(text) != 0 {
			todo.Text = text
		}

		err = db.updateTodo(todo)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to update TODO #%d: %s", id, err.Error()), 10)
		}

		return nil
	},
}
