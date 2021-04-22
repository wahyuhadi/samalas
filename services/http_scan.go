package services

import (
	"errors"
	"fmt"
	"net/http"
	loggers "samalas/logger"
	"sync"
)

var (
	list = []string{
		".env",
		".env.save",
		".env.example",
		".env.sample",
		".env.production",
		".git/config",
		"docker-compose.yml",
		"storage/logs/laravel.log",
	}
)

func (t *Target) simple_brute_dir() error {

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

		if err != nil {
			msg := fmt.Sprintf("error when get the data %s  | %s", target, err.Error())
			loggers.SetLogger("warning", msg)
			return nil
		}
		// -- only show the 200 OK
		if resp.StatusCode == http.StatusOK {
			fmt.Println(GREEN, "+ [HTTP DIR] Posible Found : ", target, RESET)
		}

		defer resp.Body.Close()
	}

	return nil
}

func (t *Target) simple_brute_dir_schema() error {

	for _, items := range list {
		client := &http.Client{
			Timeout: t.Timeout, // - hope is enough
		}

		// -- dont follow the redirect page
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		}
		target := fmt.Sprintf("%s/%s", t.Ip, items)
		req, _ := http.NewRequest("GET", target, nil)
		resp, err := client.Do(req)
		if err != nil {
			msg := fmt.Sprintf("error when get the data %s  | %s", target, err.Error())
			loggers.SetLogger("warning", msg)
			return nil
		}
		// -- only show the 200 OK
		if resp.StatusCode == http.StatusOK {
			fmt.Println(GREEN, "+ [HTTP DIR] Posible Found : ", target, RESET)
		}

		defer resp.Body.Close()

	}

	return nil
}
func (t *Target) isHttp(wg *sync.WaitGroup, withSchema bool) {
	defer wg.Done()

	if t.Http {
		if withSchema {
			// -- make http scheme false

			msg := fmt.Sprintf("Do brute directory force   on IP : %s", t.Ip)
			loggers.SetLogger("info", msg)
			t.simple_brute_dir_schema()
		} else {
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

}
