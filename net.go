/*
 * @brief 网络常用操作
 */

package encapsutils

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// PingAddress probe ip or address can be connected. NOTE: github.com/go-ping/ping must with root permitted, so use command.
func PingAddress(address string, waitTime time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ping", "-c", "1", address)
	_, err := cmd.CombinedOutput()
	return err == nil
}

// ByteToIpv6String convert ipv6 format byte byte to string.
func ByteToIpv6String(x [6]byte) string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", int(x[0]), int(x[1]), int(x[2]), int(x[3]), int(x[4]), int(x[5]))
}

// Uint32ToIpv4 convert ipv4 format uint32 to string.
func Uint32ToIpv4(ipv4 uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		uint32(ipv4&(255<<24)>>24),
		uint32(ipv4&(255<<16)>>16),
		uint32(ipv4&(255<<8)>>8),
		uint32(ipv4&(255<<0)>>0))
}

// Ipv4ToUint32 convert ipv4 format string to uint32.
func Ipv4ToUint32(ipv4 string) uint32 {
	ps := strings.Split(ipv4, ".")
	if len(ps) != 4 {
		return 0
	}

	var i uint32
	for index := 0; index < len(ps); index++ {
		tmp, err := strconv.ParseUint(ps[index], 10, 32)
		if err != nil {
			return 0
		}
		i = i | uint32(tmp<<(8*(3-index)))
	}
	return i
}
