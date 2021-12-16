package encapsutils

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestPingIPv4(t *testing.T) {
	convey.Convey("ping 127.0.0.1 right address", t, func() {
		convey.So(PingIPV4("127.0.0.1", time.Millisecond*50), convey.ShouldBeTrue)
	})

	convey.Convey("ping 1.2.3.4 error adress", t, func() {
		convey.So(PingIPV4("1.2.3.4", time.Millisecond*50), convey.ShouldBeFalse)
	})
}
