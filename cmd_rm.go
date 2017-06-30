package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

var commandRemove = cli.Command{
	Name:  "rm",
	Usage: "remove a TODO item",

	Action: func(c *cli.Context) error {
		arg := c.Args().Get(0)
		id, err := strconv.Atoi(arg)
		if err != nil {
			return cli.NewExitError("Please provide the TODO's ID as the argument", 4)
		}

		err = db.rmTodo(id)
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to remove TODO #%d", id), 5)
		}

		fmt.Printf("TODO #%d removed\n", id)
		return nil
	},
	Category: "TODOs",
}
