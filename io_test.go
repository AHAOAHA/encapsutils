package utils

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestPathExists(t *testing.T) {
	convey.Convey("probe true dir and file", t, func() {
		convey.So(PathExists("/"), convey.ShouldBeTrue)
		convey.So(PathExists("/bin/ls"), convey.ShouldBeTrue)
	})

	convey.Convey("probe false dir and file", t, func() {
		convey.So(PathExists("/ahaoo"), convey.ShouldBeFalse)
	})
}

func TestIsDirAndIsFile(t *testing.T) {
	convey.Convey("probe dir", t, func() {
		convey.So(IsDir("/"), convey.ShouldBeTrue)
		convey.So(IsDir("/bin/ls"), convey.ShouldBeFalse)
	})

	convey.Convey("probe file", t, func() {
		convey.So(IsFile("/"), convey.ShouldBeFalse)
		convey.So(IsFile("/bin/ls"), convey.ShouldBeTrue)
	})
}
