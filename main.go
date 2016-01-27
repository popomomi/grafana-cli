package main

import (
	"github.com/codegangsta/cli"
	commands "github.com/grafana/grafana-cli/commands"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "Grafana cli"
	app.Author = "raintank"
	app.Email = "https://github.com/grafana/grafana"
	app.Commands = commands.Commands

	app.Run(os.Args)
}
