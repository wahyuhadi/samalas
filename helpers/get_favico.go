package helpers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var Pattern string = "^(shortcut|icon|shortcut icon)$"
var DefaultFavicon string = "/favicon.ico"
var DefaultTag string = `<link rel="shorcut icon" href="/favicon.ico">`

type Scraper struct {
	URL      *url.URL
	LastURL  *url.URL
	LinkNode *html.Node
}

func NewScraper(u string) *Scraper {
	r, err := url.Parse(u)
	if err != nil {
		return nil
	}

	return &Scraper{URL: r}
}

func (s *Scraper) Favicon() (*url.URL, error) {
	b, err := s.get()
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	s.parse(doc)

	if s.LinkNode == nil {
		exist, err := s.head()
		if err != nil {
			return nil, err
		}

		if exist {
			dt := strings.NewReader(DefaultTag)
			dn, _ := html.Parse(dt)
			s.LinkNode = dn
		}
	}

	f := s.attr(s.LinkNode, "href")
	v, err := url.Parse(f)
	if err != nil {
		return nil, err
	}

	if v.IsAbs() {
		return v, nil
	}

	u, err := url.Parse(s.LastURL.String())
	if err != nil {
		return nil, err
	}

	r := u.ResolveReference(v)
	return r, nil
}

func (s *Scraper) parse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "link" {
		for _, a := range n.Attr {
			if a.Key == "rel" {
				match, _ := regexp.MatchString(Pattern, a.Val)
				if match == true {
					s.LinkNode = n
				}
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.parse(c)
	}
}

func (s *Scraper) attr(n *html.Node, key string) string {
	if n == nil {
		return ""
	}

	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}

func (s *Scraper) get() ([]byte, error) {
	res, err := http.Get(s.URL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// save last URL
	s.LastURL = res.Request.URL

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Scraper) head() (bool, error) {
	f := s.LastURL
	f.Path = DefaultFavicon
	f.RawQuery = ""

	res, err := http.Head(f.String())
	if err != nil {
		return false, err
	}

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
