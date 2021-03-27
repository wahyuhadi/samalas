package services

import (
	"fmt"
	"net"
	"time"
)

func (t *Target) ScanPort(port int, timeout time.Duration) bool {
	// fmt.Println(t.Ip)
	target := fmt.Sprintf("%s:%d", t.Ip, port)
	// fmt.Println(target)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
