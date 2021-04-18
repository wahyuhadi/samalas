package services

import (
	"errors"
	"fmt"
	"net/http"
	loggers "samalas/logger"
	"sync"
)

func (t *Target) simple_brute_dir() error {
	list := []string{
		".env",
		".env.example",
		".env.sample",
		".env.production",
		".git/config",
		"docker-compose.yml",
	}

	for _, items := range list {
		client := &http.Client{
			Timeout: t.Timeout, // - hope is enough
		}

		// -- dont follow the redirect page
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		}

		target := fmt.Sprintf("http://%s/%s", t.Ip, items)
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
		msg := fmt.Sprintf("Scan http port on IP : %s", t.Ip)
		loggers.SetLogger("info", msg)
		if is_http {
			msg = fmt.Sprintf("Do brute directory force   on IP : %s", t.Ip)
			loggers.SetLogger("info", msg)
			t.simple_brute_dir()
		}
	}
}
