package subdomain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	SECURITYTRAILS_BASE string = "https://api.securitytrails.com/v1"
)

type SectrailsList struct {
	Subdomains []string      `json:"subdomains"`
	Meta       SectrailsMeta `json:"meta"`
	Endpoint   string        `json:"endpoint"`
}

type SectrailsMeta struct {
	LimitReached bool
}

func ParseSecuritytrails(domainName string) []Domain {
	key := getSecurityttailsApiKey()
	domains := []Domain{}

	url := SECURITYTRAILS_BASE + "/domain/" + domainName + "/subdomains?children_only=false"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKEY", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result SectrailsList
	json.Unmarshal([]byte(body), &result)

	for _, domain := range result.Subdomains {
		rawDomain := domain + domainName

		domains = append(domains, Domain{DomainName: rawDomain, IpAddr: ""})
	}

	return domains
}

func getSecurityttailsApiKey() string {
	key := os.Getenv("SECURITYTRAILS_API_KEY")

	fmt.Println(key)

	if key == "" {
		panic("Please set SECURITYTRAILS_API_KEY on os environment")
	}

	return key
}
