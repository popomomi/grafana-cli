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
			err := lsCommand(commandLine)
			So(err, ShouldNotBeNil)
		})
	})
}
