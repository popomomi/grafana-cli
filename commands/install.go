package commands

import (
	"errors"
	"fmt"
	"io"
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

	//unzip and feast upon this great plugin!

	return err
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
