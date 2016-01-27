package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func listCommand(c *cli.Context) {
	plugin, err := listAllPlugins()

	if err != nil {
		fmt.Println("cannot find your plugin %v", err)
	}

	for _, i := range plugin.Plugins {
		fmt.Printf("id: %v version:%v\n", i.Id, i.Version)
	}
}
