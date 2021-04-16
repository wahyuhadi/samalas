package subdomain

// - handler subdomain
func HandlerSubdomain(domain string) []SubDomain {
	subDomainList := []SubDomain{}

	hackerTarget := parseHackerTarget(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: hackerTarget, Source: "HackerTarget"})

	securitytrails := ParseSecuritytrails(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: securitytrails, Source: "Securitytrails"})

	return subDomainList
}
