package xhtml_test

import (
	"net/url"
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
)

func TestAbsolutizeURL(t *testing.T) {
	testcases := []struct {
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
	for _, tc := range testcases {
		be := assert.FailsNow(t)
		n := xhtml.New("div")
		be.NilError(xhtml.SetInnerHTML(n, tc.in))
		xhtml.AbsolutizeURLs(n, u)
		be.Equal(xhtml.InnerHTML(n), tc.out)
	}
}
