package commands

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/grafana/grafana-cli/pkg/log"
	services "github.com/grafana/grafana-cli/pkg/services"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
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

	pluginToInstall := c.Args().First()

	return InstallPlugin(pluginToInstall, pluginFolder)
}

func InstallPlugin(pluginName, pluginFolder string) error {
	plugin, err := services.GetPlugin(pluginName)
	if err != nil {
		return err
	}

	downloadUrl := plugin.Url + "/archive/" + plugin.Commit + ".zip"

	log.Infof("installing %v\n", plugin.Id)
	log.Infof("from url: %v\n", downloadUrl)
	log.Infof("on commit: %v\n", plugin.Commit)
	log.Infof("into: %v\n", pluginFolder)

	return downloadFile(plugin.Id, pluginFolder, downloadUrl)
}

func RemoveGitBuildFromname(pluginname, filename string) string {
	r := regexp.MustCompile("^[a-zA-Z0-9_.-]*/")
	res := r.ReplaceAllString(filename, pluginname+"/")
	return res
}

func downloadFile(pluginName, filepath, url string) (err error) {
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
		newfile := path.Join(filepath, RemoveGitBuildFromname(pluginName, zf.Name))

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
