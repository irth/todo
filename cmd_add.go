package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

var commandAdd = cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "create a TODO item",
	UsageText: "todo add todo_text...",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "category, c",
			Value: "default",
			Usage: "the category for the TODO",
		},
	},
	Category: "TODOs",

	Action: func(c *cli.Context) error {
		task := strings.TrimSpace(strings.Join(c.Args(), " "))
		if len(task) == 0 {
			return cli.NewExitError("Provide the task description!", 42)
		}
		todo := NewTodo(task)
		if c.IsSet("category") {
			todo.Category = c.String("category")
		}
		db.addTodo(todo)
		fmt.Println("added task", task)
		return nil
	},
}
