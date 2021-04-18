package main

import (
	"flag"
	"fmt"
	"samalas/services"
	"samalas/subdomain"

	ipgen "github.com/wahyuhadi/go-ipgen"
)

var (
	subnet_ip = flag.String("subnet", "", "Scan for subnet ip Example 10.1.1.1/23")
	domain    = flag.String("domain", "", "Scan subdomain")

	ip = flag.String("ip", "", "Scan for  ip Example 10.1.1.1")
)

func main() {
	flag.Parse()

	// -- scan for subnet
	if *subnet_ip != "" {
		ip := ipgen.IpAddressGen(*subnet_ip)
		fmt.Println("[+] Run scanning ..")
		for i := 0; i < len(ip); i++ {
			ips := ip[i]
			services.Init(ips)
		}

	}

	// -- scan for subdomain
	if *domain != "" {
		subDomainList := subdomain.HandlerSubdomain(*domain)
		for _, subDomain := range subDomainList {
			for _, domain := range subDomain.GetAll() {
				// -- validate if domain name not null
				if domain.DomainName != "" {
					services.Init(domain.DomainName)
				}
				// -- validate if ip address not null
				if domain.IpAddr != "" {
					services.Init(domain.IpAddr)
				}
			}
		}
	}

	// -- scan for single IP
	if *ip != "" {
		services.Init(*ip)
	}

}
