package subdomain

type SubDomain struct {
	Domains []Domain
	Source  string
}

type Domain struct {
	IpAddr     string
	DomainName string
}

// -- handling validate if source need API - KEY
type ControllerScapper struct {
	// -- security trails need api key
	// - SECURITYTRAILS_API_KEY
	Securitytrails bool
}

func (subDomain SubDomain) GetAll() []Domain {
	domains := []Domain{}

	for _, domain := range subDomain.Domains {
		domains = append(domains, Domain{IpAddr: domain.IpAddr, DomainName: domain.DomainName})
	}

	return domains
}
