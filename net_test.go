package utils

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestPingIPv4(t *testing.T) {
	convey.Convey("ping www.baidu.com right address", t, func() {
		convey.So(PingIPV4("www.baidu.com", time.Millisecond*10), convey.ShouldBeTrue)
	})

	convey.Convey("ping 1.2.3.4 error adress", t, func() {
		convey.So(PingIPV4("1.2.3.4", time.Millisecond*10), convey.ShouldBeFalse)
	})
}
