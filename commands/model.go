package commands

import (
	"errors"
	"github.com/franela/goreq"
)

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

func listAllPlugins() (PluginRepo, error) {
	res, _ := goreq.Request{Uri: "https://raw.githubusercontent.com/grafana/grafana-cli/master/test-data/mock-repo.json"}.Do()

	var resp PluginRepo
	err := res.Body.FromJsonTo(&resp)
	if err != nil {
		return PluginRepo{}, errors.New("Could not load plugin data")
	}

	return resp, nil
}

func getPlugin(id string) (Plugin, error) {
	resp, err := listAllPlugins()
	if err != nil {
	}

	for _, i := range resp.Plugins {
		if i.Id == id {
			return i, nil
		}
	}

	return Plugin{}, errors.New("could not find plugin named \"" + id + "\"")
}
