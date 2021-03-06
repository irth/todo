package main

import (
	"os"

	"github.com/urfave/cli"
)

var db *DB

func main() {
	db = openDb()
	app := cli.NewApp()
	app.Name = "todo"
	app.UsageText = "todo command [command options] [arguments...]"
	app.Usage = "a simple CLI TODO manager"
	app.Commands = []cli.Command{
		commandAdd,
		commandRemove,
		commandUpdate,
		commandDone,
		commandList,
		commandSchedule,
		commandDeadline,
	}
	app.Run(os.Args)
}
