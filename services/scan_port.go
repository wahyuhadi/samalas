package services

import (
	"fmt"
	"net"
	"time"
)

func (t *Target) ScanPort(port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", t.Ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
