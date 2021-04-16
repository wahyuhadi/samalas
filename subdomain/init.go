package subdomain

type SubDomain struct {
	Domains []Domain
	Source  string
}

type Domain struct {
	IpAddr     string
	DomainName string
}

func (subDomain SubDomain) GetAll() []Domain {
	domains := []Domain{}

	for _, domain := range subDomain.Domains {
		domains = append(domains, Domain{IpAddr: domain.IpAddr, DomainName: domain.DomainName})
	}

	return domains
}
