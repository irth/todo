package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var commandDeadline = cli.Command{
	Name:    "Deadline",
	Aliases: []string{"de", "dead"},
	Usage:   "schedule a TODO",
	Action: func(c *cli.Context) error {
		id, err := strconv.Atoi(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID as the argument", 4)
		}

		text := strings.TrimSpace(strings.Join(c.Args().Tail(), " "))
		if text == "-" {
			todo, err := db.setDeadline(id, nil)
			if err != nil {
				return cli.NewExitError(fmt.Sprint("Couldn't remove the deadline:", err), 7)
			}
			fmt.Printf("Removed the deadline from TODO #%d (%s).\n", id, todo.Text)
			return nil
		}

		t, err := parseDate(text)
		if err != nil {
			return cli.NewExitError(err.Error(), 6)
		}

		todo, err := db.setDeadline(id, t)
		if err != nil {
			return cli.NewExitError(fmt.Sprint("Couldn't set the deadline for the task:", err), 7)
		}
		fmt.Printf("Set the deadline %s for TODO #%d (%s)\n", t.String(), id, todo.Text)
		return nil
	},
	Category: "TODOs",
}
