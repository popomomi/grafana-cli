package commands

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/grafana/grafana-cli/pkg/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func validateInput(c CommandLine, pluginFolder string) error {
	arg := c.Args().First()
	if arg == "" {
		return errors.New("please specify plugin to install")
	}

	pluginDir := c.GlobalString("path")
	if pluginDir == "" {
		return errors.New("missing path flag")
	}

	fileinfo, err := os.Stat(pluginDir)
	if err != nil && !fileinfo.IsDir() {
		return errors.New("path is not a directory")
	}

	return nil
}

func installCommand(c CommandLine) error {
	pluginFolder := c.GlobalString("path")
	if err := validateInput(c, pluginFolder); err != nil {
		return err
	}

	plugin, err := getPlugin(c.Args().First())
	if err != nil {
		return err
	}

	downloadUrl := plugin.Url + "/archive/" + plugin.Commit + ".zip"

	log.Infof("installing %v\n", plugin.Id)
	log.Infof("from url: %v\n", downloadUrl)
	log.Infof("on commit: %v\n", plugin.Commit)
	log.Infof("into: %v\n", pluginFolder)

	return downloadFile(pluginFolder, downloadUrl)
}

func downloadFile(filepath string, url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%v", err)
	}

	r, err := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if err != nil {
		log.Errorf("%v", err)
	}
	for _, zf := range r.File {
		log.Infof("filepath: %s\n", filepath)
		log.Infof("zf.Name: %s\n", zf.Name)

		newfile := path.Join(filepath, zf.Name)

		// TODO: decide how to handle the plugin folder naming
		// Depends on how we package plugins.
		// 2016-02-09 bergquist

		if zf.FileInfo().IsDir() {
			os.Mkdir(newfile, 0777)
		} else {
			dst, err := os.Create(newfile)
			if err != nil {
				log.Errorf("%v", err)
			}
			defer dst.Close()
			src, err := zf.Open()
			if err != nil {
				log.Errorf("%v", err)
			}
			defer src.Close()

			io.Copy(dst, src)
		}
	}

	return nil
}
