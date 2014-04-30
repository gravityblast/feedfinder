package feedfinder

import (
	"bytes"
	"net/url"
	"testing"

	"code.google.com/p/go.net/html"

	assert "github.com/pilu/miniassert"
)

func TestNewFinder(t *testing.T) {
	f, err := newFinder("http://example.com")
	expectedURL, _ := url.Parse("http://example.com")

	assert.Nil(t, err)
	assert.Equal(t, expectedURL, f.URL)
}

func TestFinder_FindFromReader(t *testing.T) {
	html := `<html>
	<head>
		<link rel="alternate" type="application/atom+xml" title="Atom Feed" href="http://feeds.example.com/feed">
		<link rel="alternate" type="application/rss+xml" title="Rss Feed" href="/daily.xml" />
	</head>
	<body>
	</body>
	</html>`

	f, _ := newFinder("http://example.com")
	links, _ := f.findFromReader(bytes.NewBufferString(html))

	assert.Equal(t, 2, len(links))

	expectedURL, _ := url.Parse("http://feeds.example.com/feed")
	assert.Equal(t, "http://feeds.example.com/feed", links[0].Href)
	assert.Equal(t, expectedURL, links[0].URL)
	assert.Equal(t, "application/atom+xml", links[0].Type)
	assert.Equal(t, "Atom Feed", links[0].Title)

	expectedURL, _ = url.Parse("http://example.com/daily.xml")
	assert.Equal(t, "/daily.xml", links[1].Href)
	assert.Equal(t, expectedURL, links[1].URL)
	assert.Equal(t, "application/rss+xml", links[1].Type)
	assert.Equal(t, "Rss Feed", links[1].Title)
}

func TestNewFeedLink_WithRelativeUrl(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com")
	link := newFeedLink("application/rss+xml", "/feed", "RSS", baseURL)

	assert.Equal(t, "application/rss+xml", link.Type)
	assert.Equal(t, "/feed", link.Href)
	assert.Equal(t, "RSS", link.Title)

	expectedURL, _ := url.Parse("http://example.com/feed")
	assert.Equal(t, expectedURL, link.URL)
}

func TestNewFeedLink_WithAbsoluteUrl(t *testing.T) {
	baseURL, _ := url.Parse("http://example.com")
	link := newFeedLink("application/rss+xml", "http://absolute.example.com/feed", "RSS", baseURL)

	assert.Equal(t, "application/rss+xml", link.Type)
	assert.Equal(t, "http://absolute.example.com/feed", link.Href)
	assert.Equal(t, "RSS", link.Title)

	expectedURL, _ := url.Parse("http://absolute.example.com/feed")
	assert.Equal(t, expectedURL, link.URL)
}

func TestMakeAttributesMap(t *testing.T) {
	attrs := []html.Attribute{
		{"", "href", "/feed"},
		{"", "foo", "/bar"},
	}

	m := makeAttributesMap(attrs)
	assert.Equal(t, "/feed", m["href"])
	assert.Equal(t, "/bar", m["foo"])
}
