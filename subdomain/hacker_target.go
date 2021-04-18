// - Function to add subdomain scrapper for
// - hackertarget.com
package subdomain

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	loggers "samalas/logger"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	HACKER_TARGET string = "https://hackertarget.com/find-dns-host-records/"
)

// - parsing data  from hacker target
func ParseHackerTarget(domainName string) []Domain {
	domains := []Domain{}

	resp, err := http.PostForm(HACKER_TARGET,
		url.Values{"theinput": {domainName}, "thetest": {"hostsearch"}, "name_of_nonce_field": {"f00679fe23"}, "_wp_http_referer": {"/find-dns-host-records/"}})

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	rawSubdomains := strings.Split(doc.Find("#formResponse").First().Text(), "\n")
	if rawSubdomains[0] == "API count exceeded - Increase Quota with Membership" {
		loggers.SetLogger("warning", "hacker target reach limit.")
		// -- add "" value to array
		domains = append(domains, Domain{"", ""})
		// -- and return value
		return domains

	}

	msg := fmt.Sprintf("Found %d subdomain in HackerTarget", len(rawSubdomains))
	loggers.SetLogger("info", msg)
	for _, subDomain := range rawSubdomains {
		rawDomain := strings.Split(subDomain, ",")
		domains = append(domains, Domain{rawDomain[1], rawDomain[0]})
	}

	return domains
}
