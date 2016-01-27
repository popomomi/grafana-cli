package commands

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	{
		Name:   "install",
		Usage:  "installs stuff",
		Action: installCommand,
	}, {
		Name:  "upgrade",
		Usage: "upgrades stuff",
		Action: func(c *cli.Context) {
			println("up up och iväg!")
		},
	}, {
		Name:  "remove",
		Usage: "removes stuff",
		Action: func(c *cli.Context) {
			println("nice and tidy!")
		},
	}, {
		Name:  "upgrade",
		Usage: "upgrades stuff",
		Action: func(c *cli.Context) {
			println("up up och iväg!")
		},
	},
}
