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

func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
