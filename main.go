package main

import (
	"os"

	"github.com/urfave/cli"
)

var db *DB

func main() {
	db = openDb()
	app := cli.NewApp()
	app.Usage = "a simple CLI TODO manager"
	app.Commands = []cli.Command{
		commandAdd,
		commandDone,
		commandList,
		commandSchedule,
	}
	app.Run(os.Args)
}
