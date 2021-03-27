package services

import (
	"flag"
	"time"
)

var (
	product = flag.String("p", "all", "Scan for product only . Example : --prodcuts elastic")
)

type Target struct {
	Ip      string
	Timeout time.Duration
	Http    bool
	Elastic bool
}

func Init(ip string) {
	var t Target
	// - default scan will scan all product
	if *product == "all" {
		t.Http = true
		t.Elastic = true
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
	t.Ip = ip
	t.Timeout = 500 * time.Millisecond
	t.isHttp()
	t.isElastic()
}
