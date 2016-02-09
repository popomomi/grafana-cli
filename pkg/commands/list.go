package commands

import (
	"fmt"
)

func listCommand(c CommandLine) error {
	plugin, err := listAllPlugins()

	if err != nil {
		return err
	}

	for _, i := range plugin.Plugins {
		fmt.Printf("id: %v version:%v\n", i.Id, i.Version)
	}

	return nil
}
