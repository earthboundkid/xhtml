package xhtml_test

import (
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
)

func TestDeepEqual(t *testing.T) {
	be := assert.FailsNow(t)
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
		a := be.OK(html.Parse(strings.NewReader(tc.a)))
		b := be.OK(html.Parse(strings.NewReader(tc.b)))
		be.Equal(xhtml.DeepEqual(a, b), tc.equal)
	}
}
