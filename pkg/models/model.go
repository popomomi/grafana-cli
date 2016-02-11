package models

import (
	"os"
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

type IoUtil interface {
	Stat(path string) (os.FileInfo, error)
	RemoveAll(path string) error
	ReadDir(path string) ([]os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)
}
