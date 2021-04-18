package subdomain

import (
	"os"
	loggers "samalas/logger"
)

// - handler subdomain
func HandlerSubdomain(domain string) []SubDomain {
	// -- validate for SECURITYTRAILS_API_KEY
	controller := ControllerScapper{}
	controller.Securitytrails = true // -- default value is true
	securitytrailsKey := os.Getenv("SECURITYTRAILS_API_KEY")
	if securitytrailsKey == "" {
		// -- change default value for Securitytrails
		loggers.SetLogger("error", "Stoping scan with securitytrails, you dont have api key.")
		controller.Securitytrails = false
	}

	subDomainList := []SubDomain{}

	loggers.SetLogger("info", "Get data from hackertarget.")
	hackerTarget := ParseHackerTarget(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: hackerTarget, Source: "HackerTarget"})

	loggers.SetLogger("info", "Get data from crtsh.")
	crtsh := ParseCRTSH(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: crtsh, Source: "crt.sh"})

	loggers.SetLogger("info", "Get data from netcraft.")
	netcraft := ParseNetcraft(domain)
	subDomainList = append(subDomainList, SubDomain{Domains: netcraft, Source: "Netcraft"})

	// -- if is api key true for Securitytrails
	if controller.Securitytrails {
		loggers.SetLogger("info", "Get data from SecurityTrails.")
		securitytrails := ParseSecuritytrails(domain)
		subDomainList = append(subDomainList, SubDomain{Domains: securitytrails, Source: "Securitytrails"})
	}

	return subDomainList
}
