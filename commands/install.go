package commands

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/grafana/grafana-cli/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func installCommand(c CommandLine) error {
	arg := c.Args().First()
	if arg == "" {
		return errors.New("please specify plugin to install")
	}

	plugin, err := getPlugin(arg)

	if err != nil {
		log.Error("cannot find your plugin\n")
	}

	log.Infof("installing %v\n", plugin.Id)
	log.Infof("from url: %v\n", plugin.Url)
	log.Infof("on commit: %v\n", plugin.Commit)

	downloadUrl := plugin.Url + "/archive/" + plugin.Commit + ".zip"

	err = downloadFile("tmp/", downloadUrl)

	return err
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
		path := filepath + zf.Name
		if zf.FileInfo().IsDir() {
			os.Mkdir(path, 0777)
		} else {
			log.Infof("Extracting: %s\n", zf.Name)
			dst, err := os.Create(path)
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
