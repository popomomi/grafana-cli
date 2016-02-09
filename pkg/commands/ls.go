package commands

import (
	"encoding/json"
	"errors"
	"github.com/grafana/grafana-cli/pkg/log"
	"io/ioutil"
	"os"
	"path"
)

func lsCommand(c CommandLine) error {
	pluginDir := c.GlobalString("path")
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

	files, err := ioutil.ReadDir(pluginDir)
	for _, f := range files {

		pluginData, _ := ioutil.ReadFile(path.Join(pluginDir, f.Name(), "plugin.json"))

		res := InstalledPlugin{}
		json.Unmarshal(pluginData, &res)

		log.Infof("plugin: %s @%s \n", res.Name, res.Version)
	}

	return nil
}
