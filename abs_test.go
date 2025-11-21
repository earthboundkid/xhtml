package xhtml_test

import (
	"net/url"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/xhtml"
)

func TestAbsolutizeURL(t *testing.T) {
	testcaes := []struct {
		in, out string
	}{
		{``, ``},
		{`<a href=""></a>`, `<a href=""></a>`},
		{
			`<a href="http://world.com"></a>`,
			`<a href="http://world.com"></a>`,
		},
		{
			`<a href="/file.png"></a>`,
			`<a href="http://example.com/file.png"></a>`,
		},
		{
			`<a href="file.png"></a>`,
			`<a href="http://example.com/1/file.png"></a>`,
		},
		{
			`<link href="file.css"/>`,
			`<link href="http://example.com/1/file.css"/>`,
		},
		{
			`<img href="file.css"/>`,
			`<img href="file.css"/>`,
		},
		{
			`<img src="file.css"/>`,
			`<img src="http://example.com/1/file.css"/>`,
		},
	}
	u, _ := url.Parse("http://example.com/1/")
	for _, tc := range testcaes {
		n := xhtml.New("div")
		be.NilErr(t, xhtml.SetInnerHTML(n, tc.in))
		xhtml.AbsolutizeURLs(n, u)
		be.Equal(t, tc.out, xhtml.InnerHTML(n))
	}
}
