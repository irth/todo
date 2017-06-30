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
