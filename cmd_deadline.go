package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var commandDeadline = cli.Command{
	Name:    "deadline",
	Aliases: []string{"de", "dead"},
	Usage:   "add a deadline to a TODO",
	Description: `Adds or removes a deadline to the todo.
	 Supports formats supported by https://github.com/bcampbell/fuzzytime.
	 Pass a hyphen (-) instead of a date to remove a deadline.
	 Pass just the ID to show the current deadline.`,
	ArgsUsage: "todoID deadline_date...",
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

		if len(text) == 0 {
			todo, err := db.getTodo(id)
			if err != nil {
				return cli.NewExitError("Couldn't find the TODO.", 8)
			}
			if todo.Deadline == nil {
				fmt.Printf("No deadline set for TODO #%d (%s)\n", id, todo.Text)
				return nil
			}
			fmt.Println(todo.Deadline.Format("2006-01-02 15:04"))
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
