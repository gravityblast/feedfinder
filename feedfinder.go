/*
Package feedfinder - autodiscovers feed url on web pages.
*/
package feedfinder

import (
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/html"
)

// A FeedLink represents the link tag with it's main attributes, Type, Href, Title and URL.
// The URL is always absolute, and it's based on the page's URL if the Href attribute is relative.
type FeedLink struct {
	Type  string
	Href  string
	Title string
	URL   *url.URL
}

func newFeedLink(_type, href, title string, baseURL *url.URL) *FeedLink {
	link := &FeedLink{
		Type:  _type,
		Href:  href,
		Title: title,
	}

	link.URL, _ = url.Parse(link.Href)
	if !link.URL.IsAbs() {
		link.URL.Scheme = baseURL.Scheme
		link.URL.Host = baseURL.Host
		link.URL.RawQuery = baseURL.RawQuery
	}

	return link
}

func makeAttributesMap(attrs []html.Attribute) map[string]string {
	m := make(map[string]string)

	for _, a := range attrs {
		m[a.Key] = a.Val
	}

	return m
}

type finder struct {
	URL *url.URL
}

func newFinder(rawurl string) (*finder, error) {
	var err error
	f := &finder{}
	f.URL, err = url.Parse(rawurl)

	return f, err
}

func (f *finder) findFromReader(r io.Reader) ([]*FeedLink, error) {
	var links []*FeedLink

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return links, err
	}

	doc.Find(`html>head>link[rel="alternate"]`).Each(func(i int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			a := makeAttributesMap(n.Attr)
			link := newFeedLink(a["type"], a["href"], a["title"], f.URL)
			links = append(links, link)
		}
	})

	return links, nil
}

func (f *finder) find() ([]*FeedLink, error) {
	var links []*FeedLink

	resp, err := http.Get(f.URL.String())
	if err != nil {
		return links, err
	}

	defer resp.Body.Close()

	return f.findFromReader(resp.Body)
}

// Find searches for feed links on rawurl
func Find(rawurl string) ([]*FeedLink, error) {
	var links []*FeedLink

	f, err := newFinder(rawurl)
	if err != nil {
		return links, err
	}

	return f.find()
}
