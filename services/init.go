package services

import (
	"flag"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
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
	// - will perform scanning in http service like minidirbuster
	Http bool
	// - Is Elastic scan ->
	Elastic bool
	// - Is Redis scan
	Redis bool
}

func Init(ip string, withSchema bool) {
	if ip != "" {
		var t Target
		// - default scan will scan all product
		if *product == "all" {
			// -- only scann  http dir  for this services

			if withSchema {
				t.Http = true
			} else {
				t.Http = true
				t.Elastic = true
				t.Redis = true
			}
		}

		// - spesific scan for products
		// - scan for elastic -p elastic
		if *product == "elastic" {
			// -- cannot scann url for this services
			if withSchema {
				t.Elastic = false
				logrus.Error("[+] Cannot scann with mode list in elastic")
				os.Exit(1)
			} else {
				t.Elastic = true
			}

		}

		// - scan for http services -p http
		if *product == "http" {
			t.Http = true
		}

		// - scan for redis services -p redis
		if *product == "redis" {
			// -- cannot scann url for this services
			if withSchema {
				t.Redis = false
				logrus.Error("[+] Cannot scann with mode list in redis")
				os.Exit(1)
			} else {
				t.Redis = true
			}
		}

		t.Ip = ip

		// - handling for timeout 5s
		t.Timeout = 2000 * time.Millisecond

		// - for go rutine

		runtime.GOMAXPROCS(2)

		var wg sync.WaitGroup
		wg.Add(3)
		t.isHttp(&wg, withSchema)
		t.isElastic(&wg)
		t.isRedis(&wg)
		wg.Wait()
	}
}
