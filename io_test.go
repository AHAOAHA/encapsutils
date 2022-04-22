package encapsutils

import (
	"io/ioutil"
	"os"
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

func TestCreateDirIfNotExist(t *testing.T) {
	convey.Convey("create dir if not exist", t, func() {
		convey.So(MustCreateDir("/", os.ModePerm), convey.ShouldBeNil)
		convey.So(MustCreateDir("./.ahaoaha", os.ModePerm), convey.ShouldBeNil)
	})
}

func TestMustSaveToFileAndMustAppendToFile(t *testing.T) {
	convey.Convey("file not exist whth save", t, func() {
		convey.So(MustSaveToFile([]byte("ahaoo"), "./name.txt"), convey.ShouldBeNil)
		binary, err := ioutil.ReadFile("./name.txt")
		convey.So(err, convey.ShouldBeNil)
		convey.So(binary, convey.ShouldResemble, []byte("ahaoo"))
	})

	if err := os.Remove("./name.txt"); err != nil {
		t.Fatalf(err.Error())
	}

	convey.Convey("file not exist whth append", t, func() {
		convey.So(MustAppendToFile([]byte("ahaoo"), "./name.txt"), convey.ShouldBeNil)
		binary, err := ioutil.ReadFile("./name.txt")
		convey.So(err, convey.ShouldBeNil)
		convey.So(binary, convey.ShouldResemble, []byte("ahaoo"))
		convey.So(MustAppendToFile([]byte(" zhang"), "./name.txt"), convey.ShouldBeNil)
		binary, err = ioutil.ReadFile("./name.txt")
		convey.So(err, convey.ShouldBeNil)
		convey.So(binary, convey.ShouldResemble, []byte("ahaoo zhang"))
	})
}
