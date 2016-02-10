package commands

import (
	"github.com/grafana/grafana-cli/pkg/log"
)

func listremoteCommand(c CommandLine) error {
	plugin, err := listAllPlugins()

	if err != nil {
		return err
	}

	for _, i := range plugin.Plugins {
		log.Infof("id: %v version:%v\n", i.Id, i.Version)
	}

	return nil
}
