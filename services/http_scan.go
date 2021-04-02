package services

import (
	"fmt"
	"net/http"
	"sync"
)

func (t *Target) simple_brute_dir() error {
	list := []string{
		".env",
		".git/config",
	}

	for _, items := range list {
		target := "http://" + t.Ip + "/" + items
		resp, err := http.Get(target)
		if err != nil {
			return err
		}

		if resp.StatusCode == 302 {
			return nil
		}
		if resp.StatusCode == 200 {
			fmt.Println("[+] please check manual ", target)
		}

	}

	return nil
}

func (t *Target) isHttp(wg *sync.WaitGroup) {
	defer wg.Done()
	if t.Http {
		is_http := t.ScanPort(80, t.Timeout)
		if is_http {
			t.simple_brute_dir()
		}
	}
}
