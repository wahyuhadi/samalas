package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"samalas/services"
	"samalas/subdomain"

	"github.com/sirupsen/logrus"
	ipgen "github.com/wahyuhadi/go-ipgen"
)

var (
	subnet_ip = flag.String("subnet", "", "Scan for subnet ip Example 10.1.1.1/23")
	domain    = flag.String("domain", "", "Scan subdomain")
	list      = flag.String("list", "", "domain / subdomain list scan ")
	ip        = flag.String("ip", "", "Scan for  ip Example 10.1.1.1")
)

func main() {
	flag.Parse()

	// -- scan for subnet
	if *subnet_ip != "" {
		ip := ipgen.IpAddressGen(*subnet_ip)
		fmt.Println("[+] Run scanning ..")
		for i := 0; i < len(ip); i++ {
			ips := ip[i]
			services.Init(ips, false)
		}

	}

	// -- scan for subdomain
	if *domain != "" {
		subDomainList := subdomain.HandlerSubdomain(*domain)
		for _, subDomain := range subDomainList {
			for _, domain := range subDomain.GetAll() {
				// -- validate if domain name not null
				if domain.DomainName != "" {
					services.Init(domain.DomainName, false)
				}
				// -- validate if ip address not null
				if domain.IpAddr != "" {
					services.Init(domain.IpAddr, false)
				}
			}
		}
	}

	// -- scan for single IP
	if *ip != "" {
		services.Init(*ip, false)
	}

	// -- for scanning list domain from file
	// -- schame is true (https or http)
	// -- default is false
	if *list != "" {
		f, err := os.Open(*list)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {

			services.Init(scanner.Text(), true)
		}

		if err := scanner.Err(); err != nil {
			logrus.Error("[!] Error reading files.")
		}

	}

}
