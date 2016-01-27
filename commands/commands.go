package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
)

type CommandLine interface {
	ShowHelp()
	ShowVersion()
	Application() *cli.App
	Args() cli.Args
	Bool(name string) bool
	Int(name string) int
	String(name string) string
	StringSlice(name string) []string
	GlobalString(name string) string
	FlagNames() (names []string)
	Generic(name string) interface{}
}

type contextCommandLine struct {
	*cli.Context
}

func (c *contextCommandLine) ShowHelp() {
	cli.ShowCommandHelp(c.Context, c.Command.Name)
}

func (c *contextCommandLine) ShowVersion() {
	cli.ShowVersion(c.Context)
}

func (c *contextCommandLine) Application() *cli.App {
	return c.App
}

func runCommand(command func(commandLine CommandLine) error) func(context *cli.Context) {
	return func(context *cli.Context) {

		//validate tokens etc

		if err := command(&contextCommandLine{context}); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

var Commands = []cli.Command{
	{
		Name:   "install",
		Usage:  "installs stuff",
		Action: runCommand(installCommand),
	}, {
		Name:   "list",
		Usage:  "list available plugins",
		Action: runCommand(listCommand),
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
