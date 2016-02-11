package commands

import (
	"errors"
	"github.com/grafana/grafana-cli/pkg/log"
	services "github.com/grafana/grafana-cli/pkg/services"
)

func validateCommand(pluginDir string, ioutil IoUtil) error {

	if pluginDir == "" {
		return errors.New("missing path flag")
	}

	log.Debug("plugindir: " + pluginDir + "\n")
	pluginDirInfo, err := ioutil.Stat(pluginDir)

	if err != nil {
		return errors.New("missing path flag")
	}

	if pluginDirInfo.IsDir() == false {
		return errors.New("plugin path is not a directory")
	}

	return nil

}

func lsCommand(c CommandLine, ioutil IoUtil) error {
	pluginDir := c.GlobalString("path")
	if err := validateCommand(pluginDir, ioutil); err != nil {
		return err
	}

	plugins := services.GetLocalPlugins(pluginDir)

	for _, plugin := range plugins {
		log.Infof("plugin: %s @ %s \n", plugin.Name, plugin.Info.Version)
	}

	return nil
}
