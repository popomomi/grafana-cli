package commands

import (
	"github.com/grafana/grafana-cli/pkg/log"
	services "github.com/grafana/grafana-cli/pkg/services"
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

	localPlugins := services.GetLocalPlugins(pluginDir)

	remotePlugins, err := services.ListAllPlugins()

	if err != nil {
		return err
	}

	pluginsToUpgrade := make([]services.InstalledPlugin, 0)

	for _, localPlugin := range localPlugins {
		for _, remotePlugin := range remotePlugins.Plugins {
			if localPlugin.Id == remotePlugin.Id {
				if ShouldUpgrade(localPlugin.Info.Version, remotePlugin.Version) {
					pluginsToUpgrade = append(pluginsToUpgrade, localPlugin)
				}
			}
		}
	}

	for _, p := range pluginsToUpgrade {
		log.Infof("lets upgrade %v \n", p)

		//remote local plugin p
		//install plugin p
	}
	// download new plugins

	return nil
}
