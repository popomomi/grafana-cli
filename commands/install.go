package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"net/http"
	"os"
)

func installCommand(c *cli.Context) {
	arg := c.Args().First()
	if arg == "" {
		fmt.Println("please specify plugin to install")
		return
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

	if err != nil {
		fmt.Printf("%v", err)
	}
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
