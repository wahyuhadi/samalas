package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/olivere/elastic"
)

func (t *Target) isElastic(wg *sync.WaitGroup) {
	defer wg.Done()
	if t.Elastic {

		is_elasic := t.ScanPort(9200, t.Timeout)
		// - if elastic open
		if is_elasic {
			ip := "http://" + t.Ip + ":9200"
			if is_elastic(ip) {
				fmt.Println("[+] port 9200 found : ", t.Ip)
			}
		}
	}
}

func is_elastic(setURL string) bool {
	var sec time.Duration = 5

	// Convert port integer to a string
	timeOut := sec * time.Second

	// Instantiate a client instance of the elastic library
	_, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(setURL),
		elastic.SetHealthcheckInterval(timeOut), // quit trying after 5 seconds
	)

	if err != nil {
		return false // - if not found return false
	}

	return true
}
