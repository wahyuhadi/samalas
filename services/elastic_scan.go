package services

import (
	"fmt"
	loggers "samalas/logger"
	"sync"
	"time"

	"github.com/olivere/elastic"
)

func (t *Target) isElastic(wg *sync.WaitGroup) {
	defer wg.Done()
	if t.Elastic {

		is_elasic := t.ScanPort(9200, t.Timeout)
		msg := fmt.Sprintf("Scanning IP :%s for elastic services.", t.Ip)
		loggers.SetLogger("info", msg)
		// - if elastic open
		if is_elasic {
			ip := fmt.Sprintf("http://%s:9200", t.Ip) // - elastic link
			if check_elastic(ip) {
				fmt.Println("[+] port 9200 found : ", t.Ip)
			}
		}
	}
}

func check_elastic(setURL string) bool {
	var sec time.Duration = 5
	msg := fmt.Sprintf("Validate IP :%s for elastic services.", setURL)
	loggers.SetLogger("info", msg)
	// Convert port integer to a string
	timeOut := sec * time.Second

	// Instantiate a client instance of the elastic library
	_, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(setURL),
		elastic.SetHealthcheckInterval(timeOut), // quit trying after 5 seconds
	)

	if err != nil {
		msg = fmt.Sprintf("Disable elastic in IP :%s , with error %s", setURL, err)
		loggers.SetLogger("info", msg)
		return false // - if not found return false
	}
	msg = fmt.Sprintf("Yeaaayy, Found elastic services in IP :%s", setURL)
	loggers.SetLogger("info", msg)
	return true
}
