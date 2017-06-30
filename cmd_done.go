package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

var commandDone = cli.Command{
	Name:    "done",
	Aliases: []string{"d"},
	Usage:   "mark a TODO item as done",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "undo",
			Usage: "mark as *not* done",
		},
	},
	Action: func(c *cli.Context) error {
		arg := c.Args().Get(0)
		id, err := strconv.Atoi(arg)
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID as the argument", 4)
		}
		todo, err := db.markTodoAsDone(id, !c.Bool("undo"))

		msg := "done"
		if c.Bool("undo") {
			msg = "not done"
		}

		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to mark TODO #%d as %s", id, msg), 5)
		}

		fmt.Printf("TODO #%d (%s) marked as %s.\n", id, todo.Text, msg)
		return nil
	},
	Category: "TODOs",
}
