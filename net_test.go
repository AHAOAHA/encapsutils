package encapsutils

import (
	"math"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestPingAddress(t *testing.T) {
	convey.Convey("ping 127.0.0.1 right address", t, func() {
		convey.So(PingAddress("127.0.0.1", time.Millisecond*50), convey.ShouldBeTrue)
	})

	convey.Convey("ping 1.2.3.4 error adress", t, func() {
		convey.So(PingAddress("1.2.3.4", time.Millisecond*50), convey.ShouldBeFalse)
	})
}

func TestIpConvert(t *testing.T) {
	convey.Convey("uint32 max to ipv4", t, func() {
		convey.So(Uint32ToIpv4(math.MaxUint32), convey.ShouldEqual, "255.255.255.255")
	})

	convey.Convey("255.255.255.255 to ipv4 uint32", t, func() {
		convey.So(Ipv4ToUint32("255.255.255.255"), convey.ShouldEqual, math.MaxUint32)
	})

}
