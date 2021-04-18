package subdomain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	loggers "samalas/logger"
)

const (
	NETCRAFT_BASE_URL string = "https://searchdns.netcraft.com/?restriction=site+contains&host="
)

func ParseNetcraft(domainName string) []Domain {
	domains := []Domain{}

	resp, err := http.Get(CRTSH_BASE_URL + domainName)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`\w+.[.]` + domainName)
	rawSubDomain := r.FindAllString(string(body), -1)
	msg := fmt.Sprintf("Found %d subdomain in Netcraft", len(rawSubDomain))
	loggers.SetLogger("info", msg)
	for _, subdomain := range removeDuplicateValues(rawSubDomain) {
		domains = append(domains, Domain{DomainName: subdomain, IpAddr: ""})
	}

	return domains
}
