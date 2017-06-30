package main

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func getPrintableDate(d *time.Time) string {
	if d != nil {
		return d.Format("2006-01-02 15:04")
	}
	return "-"
}

func getPrintableTodo(todo Todo) []string {
	newTodo := []string{}

	newTodo = append(newTodo, strconv.Itoa(todo.ID))
	if todo.Done {
		newTodo = append(newTodo, "✔")
	} else {
		newTodo = append(newTodo, " ")
	}
	newTodo = append(newTodo, todo.Text)
	newTodo = append(newTodo, getPrintableDate(todo.Date))
	newTodo = append(newTodo, getPrintableDate(todo.Deadline))
	newTodo = append(newTodo, todo.Category)

	return newTodo
}

var commandList = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "list TODO items",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "show all TODOs, even those that are marked as done",
		},
		cli.StringFlag{
			Name:  "category, c",
			Usage: "if specified, only todos from this category will be shown",
		},
	},
	Action: func(c *cli.Context) error {
		var category *string
		if c.IsSet("category") {
			categoryString := c.String("category")
			category = &categoryString
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "?", "Todo", "Due", "Deadline", "Category"})
		for _, todo := range db.getTodos(c.Bool("all"), category) {
			table.Append(getPrintableTodo(todo))
		}
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.Render()
		return nil
	},
	Category: "TODOs",
}
