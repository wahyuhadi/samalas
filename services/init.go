package services

import (
	"time"
)

type Target struct {
	Ip      string
	Timeout time.Duration
}

func Init(ip string) {
	var t Target
	t.Ip = ip
	t.Timeout = 500 * time.Millisecond
	t.isHttp()
	t.isElastic()

}
