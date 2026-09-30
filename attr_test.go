package xhtml_test

import (
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
)

func TestAttr(t *testing.T) {
	be := assert.FailsNow(t)
	var n *html.Node
	be.Equal("", xhtml.Attr(n, "a"))
	n = xhtml.New("span")
	be.Equal(`<span></span>`, xhtml.OuterHTML(n))
	be.Equal("", xhtml.Attr(n, "a"))

	xhtml.SetAttr(n, "a", "b")
	be.Equal(`<span a="b"></span>`, xhtml.OuterHTML(n))
	be.Equal("b", xhtml.Attr(n, "a"))

	xhtml.SetAttr(n, "a", "c")
	be.Equal(`<span a="c"></span>`, xhtml.OuterHTML(n))
	be.Equal("c", xhtml.Attr(n, "a"))

	xhtml.SetAttr(n, "d", "e")
	be.Equal("c", xhtml.Attr(n, "a"))
	be.Equal("e", xhtml.Attr(n, "d"))
	be.Equal(`<span a="c" d="e"></span>`, xhtml.OuterHTML(n))

	xhtml.DeleteAttr(n, "a")
	be.Equal("", xhtml.Attr(n, "a"))
	be.Equal("e", xhtml.Attr(n, "d"))
	be.Equal(`<span d="e"></span>`, xhtml.OuterHTML(n))
}
