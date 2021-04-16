package subdomain

import (
	"os"
)

// - handler subdomain
func HandlerSubdomain(domain string) []SubDomain {
	// -- validate for SECURITYTRAILS_API_KEY
	controller := ControllerScapper{}
	controller.Securitytrails = true // -- default value is true
	key := os.Getenv("SECURITYTRAILS_API_KEY")
	if key == "" {
		// -- change default value for Securitytrails
		controller.Securitytrails = false
	}

	subDomainList := []SubDomain{}

	hackerTarget := parseHackerTarget(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: hackerTarget, Source: "HackerTarget"})

	// -- if is api key true for Securitytrails
	if controller.Securitytrails {
		securitytrails := ParseSecuritytrails(domain)
		subDomainList = append(subDomainList, SubDomain{Domains: securitytrails, Source: "Securitytrails"})
	}

	return subDomainList
}
