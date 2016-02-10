package commands

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFoldernameReplacement(t *testing.T) {
	Convey("path containing git commit path", t, func() {
		paths := map[string]string{
			"datasource-plugin-kairosdb-cc4a3965ef5d3eb1ae0ee4f93e9e78ec7db69e64/":                     "datasource-plugin-kairosdb/",
			"datasource-plugin-kairosdb-cc4a3965ef5d3eb1ae0ee4f93e9e78ec7db69e64/README.md":            "datasource-plugin-kairosdb/README.md",
			"datasource-plugin-kairosdb-cc4a3965ef5d3eb1ae0ee4f93e9e78ec7db69e64/partials/":            "datasource-plugin-kairosdb/partials/",
			"datasource-plugin-kairosdb-cc4a3965ef5d3eb1ae0ee4f93e9e78ec7db69e64/partials/config.html": "datasource-plugin-kairosdb/partials/config.html",
		}

		Convey("should be replaced with plugin name", func() {
			for k, v := range paths {
				So(FormatFilename("datasource-plugin-kairosdb", k), ShouldEqual, v)
			}
		})
	})
}
