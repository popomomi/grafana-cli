package commands

import (
	"testing"

	"github.com/grafana/grafana-cli/pkg/commands/commandstest"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMissingPath(t *testing.T) {
	Convey("Missing path", t, func() {
		commandLine := &commandstest.FakeCommandLine{
			CliArgs: []string{"ls"},
			GlobalFlags: &commandstest.FakeFlagger{
				Data: map[string]interface{}{
					"path": "",
				},
			},
		}

		Convey("should return error", func() {
			err := lsCommand(commandLine, &IoUtilImp{})
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Path is not a directory", t, func() {
		commandLine := &commandstest.FakeCommandLine{
			CliArgs: []string{"ls"},
			GlobalFlags: &commandstest.FakeFlagger{
				Data: map[string]interface{}{
					"path": "/var/lib/grafana/plugins",
				},
			},
		}

		util := &commandstest.FakeIoUtil{}

		Convey("should return error", func() {
			err := lsCommand(commandLine, util)
			So(err, ShouldNotBeNil)
		})
	})
}
