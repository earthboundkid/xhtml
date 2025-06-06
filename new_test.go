package xhtml_test

import (
	"strings"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestClone(t *testing.T) {
	cases := []string{
		"",
		"<p>hello, world</p>",
		`<h1><a href="http://example.com">link</a></h1><div>boo</div>`,
	}
	for _, tc := range cases {
		n, err := html.Parse(strings.NewReader(tc))
		be.NilErr(t, err)
		body := n.FirstChild.FirstChild.NextSibling
		be.Equal(t, xhtml.Select(n, xhtml.WithAtom(atom.Body)), body)

		s := xhtml.InnerHTML(body)
		be.Equal(be.Relaxed(t), tc, s)

		n2 := xhtml.Clone(n)
		body2 := n2.FirstChild.FirstChild.NextSibling
		be.Equal(t, xhtml.Select(n2, xhtml.WithAtom(atom.Body)), body2)
		be.Unequal(t, body, body2)
		be.True(t, xhtml.DeepEqual(body, body2))

		s = xhtml.InnerHTML(body2)
		be.Equal(be.Relaxed(t), tc, s)

		m := map[*html.Node]bool{}
		for c := range n.Descendants() {
			m[c] = true
		}

		for c := range n2.Descendants() {
			if m[c] {
				t.Error("duplicate node:", n)
			}
		}
	}
}
