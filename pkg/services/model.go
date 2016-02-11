package services

import (
	"encoding/json"
	"errors"
	"github.com/franela/goreq"
	"io/ioutil"
	"path"
)

type InstalledPlugin struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Info PluginInfo `json:"info"`
}

type PluginInfo struct {
	Version string `json:"version"`
	Updated string `json:"updated"`
}

type Plugin struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Commit   string `json:"commit"`
	Url      string `json:"url"`
	Version  string `json:"version"`
}

type PluginRepo struct {
	Plugins []Plugin `json:"plugins"`
	Version string   `json:"version"`
}

func ListAllPlugins() (PluginRepo, error) {
	res, _ := goreq.Request{Uri: "https://raw.githubusercontent.com/grafana/grafana-plugin-repository/master/repo.json"}.Do()

	var resp PluginRepo
	err := res.Body.FromJsonTo(&resp)
	if err != nil {
		return PluginRepo{}, errors.New("Could not load plugin data")
	}

	return resp, nil
}

func GetLocalPlugins(pluginDir string) []InstalledPlugin {
	result := make([]InstalledPlugin, 0)

	files, _ := ioutil.ReadDir(pluginDir)
	for _, f := range files {
		pluginData, _ := ioutil.ReadFile(path.Join(pluginDir, f.Name(), "plugin.json"))

		res := InstalledPlugin{}
		json.Unmarshal(pluginData, &res)

		if res.Info.Version == "" {
			res.Info.Version = "0.0.0"
		}

		if res.Id == "" {
			res.Id = res.Name
		}

		result = append(result, res)
	}

	return result
}

func GetPlugin(id string) (Plugin, error) {
	resp, err := ListAllPlugins()
	if err != nil {
	}

	for _, i := range resp.Plugins {
		if i.Id == id {
			return i, nil
		}
	}

	return Plugin{}, errors.New("could not find plugin named \"" + id + "\"")
}
