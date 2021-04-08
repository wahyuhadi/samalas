package services

import (
	"errors"
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
		client := &http.Client{}
		// -- dont follow the redirect page
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		}
		target := "http://" + t.Ip + "/" + items
		req, _ := http.NewRequest("GET", target, nil)
		resp, err := client.Do(req)

		if err == nil {
			if resp.StatusCode == http.StatusOK {
				fmt.Println("[+] Posible Found : ", target)
			}
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
