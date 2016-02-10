package commands

import (
	"github.com/codegangsta/cli"
	"github.com/grafana/grafana-cli/pkg/log"
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

		cmd := &contextCommandLine{context}
		if err := command(cmd); err != nil {
			log.Errorf("%v\n\n", err)

			cmd.ShowHelp()
		}
	}
}

var Commands = []cli.Command{
	{
		Name:   "install",
		Usage:  "installs stuff",
		Action: runCommand(installCommand),
	}, {
		Name:   "list-remote",
		Usage:  "list remote available plugins",
		Action: runCommand(listremoteCommand),
	}, {
		Name:   "upgrade",
		Usage:  "upgrades one plugin",
		Action: runCommand(upgradeCommand),
	}, {
		Name:   "upgrade-all",
		Usage:  "upgrades all your installed plugins",
		Action: runCommand(upgradeAllCommand),
	}, {
		Name:   "ls",
		Usage:  "list all installed plugins",
		Action: runCommand(lsCommand),
	}, {
		Name:   "remove",
		Usage:  "removes stuff",
		Action: runCommand(removeCommand),
	},
}
