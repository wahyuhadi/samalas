package subdomain

type SubDomain struct {
	Subdomain string
	IP        string
}

func HandlerSubdomain(domain string) []SubDomain {
	SubDomainList := []SubDomain{}

	hackerTarget := parseHackerTarget(domain)

	// - Add subdomain to array object fot hacker target
	for _, htarget := range hackerTarget {
		SubDomainList = append(SubDomainList, SubDomain{htarget.DomainName, htarget.IpAddr})
	}
	return SubDomainList
}
