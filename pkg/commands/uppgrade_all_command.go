package commands

import (
	"github.com/grafana/grafana-cli/pkg/log"
)

func upgradeAllCommand(c CommandLine) error {
	pluginDir := c.GlobalString("path")

	localPlugins := getLocalPlugins(pluginDir)

	remotePlugins, err := listAllPlugins()

	if err != nil {
		return err
	}

	for _, localPlugin := range localPlugins {
		for _, remotePlugin := range remotePlugins.Plugins {
			log.Infof("%s ==  %s\n", localPlugin.Id, remotePlugin.Id)
			if localPlugin.Id == remotePlugin.Id {
				log.Infof("\tShould I upgrade %s from %s to %s?", localPlugin.Name, localPlugin.Info.Version, remotePlugin.Version)
			}
		}
	}

	// compare version

	// download new plugins

	return nil
}
