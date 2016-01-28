package commands

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
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
		fmt.Println("cannot find your plugin")
	}

	fmt.Printf("installing %v\n", plugin.Id)
	fmt.Printf("from url: %v\n", plugin.Url)
	fmt.Printf("on commit: %v\n", plugin.Commit)

	downloadUrl := plugin.Url + "/archive/" + plugin.Commit + ".zip"
	localfileName := plugin.Id + ".zip"

	err = downloadFile(localfileName, downloadUrl)

	return err
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

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
		fmt.Println(zf.Name)

		path := "tmp/" + zf.Name
		if zf.FileInfo().IsDir() {
			os.Mkdir(path, 0777)
		} else {
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
