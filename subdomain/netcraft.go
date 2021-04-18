package subdomain

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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
	rawSubDoamin := r.FindAllString(string(body), -1)

	for _, subdomain := range removeDuplicateValues(rawSubDoamin) {
		domains = append(domains, Domain{DomainName: subdomain, IpAddr: ""})
	}

	return domains
}
