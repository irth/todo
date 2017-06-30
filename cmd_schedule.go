package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var commandSchedule = cli.Command{
	Name:    "schedule",
	Aliases: []string{"s", "due"},
	Usage:   "schedule a TODO",
	Description: `Adds or removes a due date to the todo.
	 Supports formats supported by https://github.com/bcampbell/fuzzytime.
	 Pass a hyphen (-) instead of a date to remove the due date.
	 Pass just the ID to show the current due date.`,
	ArgsUsage: "todoID due_date...",
	Action: func(c *cli.Context) error {
		id, err := strconv.Atoi(c.Args().Get(0))
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID as the argument", 4)
		}

		text := strings.TrimSpace(strings.Join(c.Args().Tail(), " "))
		if text == "-" {
			todo, err := db.setDate(id, nil)
			if err != nil {
				return cli.NewExitError(fmt.Sprint("Couldn't remove the due date:", err), 7)
			}
			fmt.Printf("Removed the due date from TODO #%d (%s).\n", id, todo.Text)
			return nil
		}

		if len(text) == 0 {
			todo, err := db.getTodo(id)
			if err != nil {
				return cli.NewExitError("Couldn't find the TODO.", 8)
			}
			if todo.Date == nil {
				fmt.Printf("No due date set for TODO #%d (%s)\n", id, todo.Text)
				return nil
			}
			fmt.Println(todo.Date.Format("2006-01-02 15:04"))
			return nil
		}

		t, err := parseDate(text)
		if err != nil {
			return cli.NewExitError(err.Error(), 6)
		}

		todo, err := db.setDate(id, t)
		if err != nil {
			return cli.NewExitError(fmt.Sprint("Couldn't schedule the task:", err), 7)
		}
		fmt.Printf("Scheduled TODO #%d (%s) for %s\n", id, todo.Text, t.String())
		return nil
	},
	Category: "TODOs",
}
