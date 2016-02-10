package commands

import (
	"github.com/grafana/grafana-cli/pkg/log"
	"github.com/hashicorp/go-version"
)

func ShouldUpgrade(installed, remote string) bool {
	installedVersion, err1 := version.NewVersion(installed)
	remoteVersion, err2 := version.NewVersion(remote)

	if err1 != nil || err2 != nil {
		return false
	}

	return installedVersion.LessThan(remoteVersion)
}

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
				var upgrade = ShouldUpgrade(localPlugin.Info.Version, remotePlugin.Version)
				log.Infof("\tShould I upgrade %s from %s to %s? %v \n\n",
					localPlugin.Name,
					localPlugin.Info.Version,
					remotePlugin.Version,
					upgrade)

			}
		}
	}

	// download new plugins

	return nil
}
