package services

import (
	"flag"
	"time"
)

var (
	product = flag.String("p", "all", "Scan for product only . Example : -p elastic")
)

type Target struct {
	// - ip a.k.a target
	Ip string
	// - timeout
	Timeout time.Duration
	// - Is Http scan -> true
	Http bool
	// - Is Elastic scan ->
	Elastic bool
	// - Is Redis scan
	Redis bool
}

func Init(ip string) {
	var t Target
	// - default scan will scan all product
	if *product == "all" {
		t.Http = true
		t.Elastic = true
		t.Redis = true
	}

	// - spesific scan for products
	// - scan for elastic
	if *product == "elastic" {
		t.Elastic = true
	}
	// - scan for http services
	if *product == "http" {
		t.Http = true
	}

	// - scan for redis services
	if *product == "redis" {
		t.Redis = true
	}

	t.Ip = ip
	t.Timeout = 500 * time.Millisecond
	t.isHttp()
	t.isElastic()
	t.isRedis()
}
