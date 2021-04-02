package main

import (
	"flag"
	"fmt"
	"samalas/services"

	ipgen "github.com/wahyuhadi/go-ipgen"
)

var (
	subnet_ip = flag.String("subnet", "", "Scan for subnet ip Example 10.1.1.1/23")
	ip        = flag.String("ip", "", "Scan for  ip Example 10.1.1.1")
)

func main() {
	flag.Parse()

	if *subnet_ip != "" {
		ip := ipgen.IpAddressGen(*subnet_ip)
		fmt.Println("[+] Run scanning ..")
		for i := 0; i < len(ip); i++ {
			ips := ip[i]
			services.Init(ips)
		}

	}

	if *ip != "" {
		services.Init(*ip)
	}

}
