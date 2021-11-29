/*
 * @brief 网络常用操作
 */

package utils

import (
	"context"
	"os/exec"
	"time"
)

// github.com/go-ping/ping must with root permitted, so use command.
func PingIPV4(ip string, waitTime time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ping", "-c", "1", ip)
	_, err := cmd.CombinedOutput()
	return err == nil
}
