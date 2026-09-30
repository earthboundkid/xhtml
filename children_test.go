package xhtml_test

import (
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html/atom"
)

func TestSetInnerHTML(t *testing.T) {
	be := assert.FailsNow(t)
	n := xhtml.New("p")
	be.NilError(xhtml.SetInnerHTML(n, "Hello, <i>World!</i>"))
	be.Equal(`<p>Hello, <i>World!</i></p>`, xhtml.OuterHTML(n))

	be.NilError(xhtml.SetInnerHTML(n, "Jello, <i>World!</i>"))
	be.Equal(`<p>Jello, <i>World!</i></p>`, xhtml.OuterHTML(n))

	n = xhtml.New("script")
	be.NilError(xhtml.SetInnerHTML(n, "let i = 1 > 2"))
	be.Equal(`<script>let i = 1 > 2</script>`, xhtml.OuterHTML(n))

	n = xhtml.New("p")
	be.NilError(xhtml.SetInnerHTML(n, "</a></b>"))
	be.Equal(`<p></p>`, xhtml.OuterHTML(n))
}

func TestUnnestChildren(t *testing.T) {
	be := assert.FailsNow(t)
	n := xhtml.New("div")
	be.NilError(xhtml.SetInnerHTML(n,
		`<a><b><i>test</i> <i>one</i> <em><i>two</i></em> </b></a>`))

	{
		clone := xhtml.Clone(n)
		i := xhtml.Select(clone, xhtml.WithAtom(atom.I))
		xhtml.UnnestChildren(i)
		be.Equal(`<a><b>test <i>one</i> <em><i>two</i></em> </b></a>`,
			xhtml.InnerHTML(clone))
	}
	{
		clone := xhtml.Clone(n)
		em := xhtml.Select(clone, xhtml.WithAtom(atom.Em))
		xhtml.UnnestChildren(em)
		be.Equal(`<a><b><i>test</i> <i>one</i> <i>two</i> </b></a>`,
			xhtml.InnerHTML(clone))
	}
	{
		clone := xhtml.Clone(n)
		a := xhtml.Select(clone, xhtml.WithAtom(atom.A))
		xhtml.UnnestChildren(a)
		be.Equal(`<b><i>test</i> <i>one</i> <em><i>two</i></em> </b>`,
			xhtml.InnerHTML(clone))
	}
	{
		clone := xhtml.Clone(n)
		b := xhtml.Select(clone, xhtml.WithAtom(atom.B))
		xhtml.UnnestChildren(b)
		be.Equal(`<a><i>test</i> <i>one</i> <em><i>two</i></em> </a>`,
			xhtml.InnerHTML(clone))
	}
	{
		clone := xhtml.Clone(n)
		for _, c := range xhtml.SelectSlice(clone, xhtml.WithAtom(atom.I)) {
			xhtml.UnnestChildren(c)
		}
		be.Equal(`<a><b>test one <em>two</em> </b></a>`,
			xhtml.InnerHTML(clone))
	}
}
