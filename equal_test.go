package xhtml_test

import (
	"strings"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
)

func TestDeepEqual(t *testing.T) {
	cases := []struct {
		a, b  string
		equal bool
	}{
		{"", "", true},
		{"<a></a>", "<b></b>", false},
		{"<p>hello, world</p>", "<p></p>", false},
		{
			`<h1><a href="http://example.com">link</a></h1><div>boo</div>`,
			`<h1><a href="http://example.com">link</a></h1><div>boo</div>`,
			true,
		},
		{
			`<h1><a href="http://example.com">link</a></h1><div>boo</div>`,
			`<h1><a href="http://example.com/">link</a></h1><div>boo</div>`,
			false,
		},
		{
			"<div><span></span><span>a</span></div>",
			"<div><span><span>a</span></span></div>",
			false,
		},
	}
	for _, tc := range cases {
		a, err := html.Parse(strings.NewReader(tc.a))
		be.NilErr(t, err)
		b, err := html.Parse(strings.NewReader(tc.b))
		be.NilErr(t, err)
		be.Equal(t, tc.equal, xhtml.DeepEqual(a, b))
	}
}
