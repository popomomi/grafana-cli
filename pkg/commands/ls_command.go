package commands

import (
	"errors"
	"github.com/grafana/grafana-cli/pkg/log"

	"os"
)

func validateCommand(pluginDir string) error {

	if pluginDir == "" {
		return errors.New("missing path flag")
	}

	log.Debug("plugindir: " + pluginDir + "\n")
	pluginDirInfo, err := os.Stat(pluginDir)

	if err != nil {
		return errors.New("missing path flag")
	}

	if pluginDirInfo.IsDir() == false {
		return errors.New("plugin path is not a directory")
	}

	return nil

}

func lsCommand(c CommandLine) error {
	pluginDir := c.GlobalString("path")
	if err := validateCommand(pluginDir); err != nil {
		return err
	}

	plugins := getLocalPlugins(pluginDir)

	for _, plugin := range plugins {
		log.Infof("plugin: %s @ %s \n", plugin.Name, plugin.Info.Version)
	}

	return nil
}
