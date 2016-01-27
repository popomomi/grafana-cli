package commands

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/franela/goreq"
)

type Plugin struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Commit   string `json:"commit"`
	Url      string `json:"url"`
	version  string `json:"version"`
}

type PluginRepo struct {
	Plugins []Plugin `json:"plugins"`
	Version string   `json:"version"`
}

func getPlugin(id string) (Plugin, error) {
	res, _ := goreq.Request{Uri: "https://raw.githubusercontent.com/grafana/grafana-cli/master/test-data/mock-repo.json"}.Do()

	var resp PluginRepo
	err := res.Body.FromJsonTo(&resp)
	fmt.Println(err)

	for _, i := range resp.Plugins {
		//fmt.Println(i.Id)
		if i.Id == id {
			return i, nil
		}
	}

	return Plugin{}, errors.New("could not find ")
}

func installCommand(c *cli.Context) {
	plugin, err := getPlugin("panel-plugin-piechart")

	if err != nil {
		fmt.Println("cannot find your plugin")
	}

	fmt.Printf("installing %v\n", plugin.Id)
	fmt.Printf("from url: %v\n", plugin.Url)
	fmt.Printf("on commit: %v\n", plugin.Commit)
}
