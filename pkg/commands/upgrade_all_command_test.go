package commands

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersionComparsion(t *testing.T) {
	Convey("Validate that version is outdated", t, func() {
		shouldUpgrade := map[string]string{
			"0.0.0": "1.0.0",
			"1.0.0": "1.0.1",
			"1.1.0": "1.1.1",
		}

		Convey("should return error", func() {
			for k, v := range shouldUpgrade {
				So(ShouldUpgrade(k, v), ShouldBeTrue)
			}
		})
	})

	Convey("Validate that version is ok", t, func() {
		shouldNotUpgrade := map[string]string{
			"2.0.0": "1.91.91",
			"3.0.0": "3.0.0",
			"2.1.1": "2.0.91",
			"x":     "1.0.0",
		}

		Convey("should return error", func() {
			for k, v := range shouldNotUpgrade {
				So(ShouldUpgrade(k, v), ShouldBeFalse)
			}
		})
	})
}
