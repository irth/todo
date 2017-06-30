package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

var commandDone = cli.Command{
	Name:      "done",
	Aliases:   []string{"d"},
	Usage:     "mark a TODO item as done",
	UsageText: "todo done -i ID [--undo]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "undo",
			Usage: "mark as *not* done",
		},
		cli.StringFlag{
			Name:  "id, i",
			Usage: "ID of the TODO that you want to mark as done",
		},
	},
	Category: "TODOs",

	Action: func(c *cli.Context) error {
		arg := c.String("id")
		id, err := strconv.Atoi(arg)
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID", 4)
		}
		todo, err := db.setDone(id, !c.Bool("undo"))

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
}
