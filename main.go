package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/grafana/grafana-cli/commands"
	"github.com/grafana/grafana-cli/log"
	"os"
)

func getGFPath() string {
	return "/var/lib/grafana" //based on your OS!
}

func main() {
	SetupLogging()

	app := cli.NewApp()
	app.Name = "Grafana cli"
	app.Author = "raintank"
	app.Email = "https://github.com/grafana/grafana"
	app.Version = "0.0.1"
	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable Verbose printing",
		},
		cli.StringFlag{
			Name:   "grafana path, p",
			Usage:  "path to the grafana installation",
			EnvVar: "GF_PATH",
			Value:  getGFPath(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorf("%v", err)
	}
}

func SetupLogging() {
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			log.SetDebug(true)
		}
	}
}

func cmdNotFound(c *cli.Context, command string) {
	fmt.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.\n",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}
