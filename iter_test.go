package xhtml_test

import (
	"fmt"
	"strings"

	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ExampleSelectAll() {
	doc, err := html.Parse(strings.NewReader(`
	<ul>
		<li><a href="https://example.com/en-us">Hello, World!</a></li>
		<li><a href="https://example.com/ja-jp">こんにちは世界！</a></li>
	</ul>`))
	if err != nil {
		panic(err)
	}
	// Find all links in a document
	// And print the link URL and text
	for link := range xhtml.SelectAll(doc, xhtml.WithAtom(atom.A)) {
		fmt.Println(xhtml.Attr(link, "href"), xhtml.TextContent(link))
	}
	// Output:
	// https://example.com/en-us Hello, World!
	// https://example.com/ja-jp こんにちは世界！
}

func ExampleClosest() {
	doc, err := html.Parse(strings.NewReader(`
	<div data-server-id="abc123">
		<button id="theButton"></button
	</div>`))
	if err != nil {
		panic(err)
	}
	// Find #theButton
	bttnEl := xhtml.Select(doc, xhtml.WithID("theButton"))
	// Find its data-server-id
	serverIDEl := xhtml.Closest(bttnEl, xhtml.WithDataset("server-id"))
	serverID := xhtml.Attr(serverIDEl, "data-server-id")
	fmt.Println(serverID)
	// Output:
	// abc123
}

func ExampleDescendantsDepth() {
	{
		s := "<div><span></span><span>a</span></div>"
		fmt.Printf("DescendantsDepth(%q)\n", s)
		doc, err := html.Parse(strings.NewReader(s))
		if err != nil {
			panic(err)
		}
		body := doc.FirstChild.FirstChild.NextSibling
		for depth, n := range xhtml.DescendantsDepth(body) {
			ntype := "element"
			if n.Type != html.ElementNode {
				ntype = "text"
			}
			fmt.Println("depth:", depth, "type:", ntype, "data:", n.Data)
		}
	}
	{
		s := "<div><span><span>a</span></span></div>"
		fmt.Printf("DescendantsDepth(%q)\n", s)
		doc, err := html.Parse(strings.NewReader(s))
		if err != nil {
			panic(err)
		}
		body := doc.FirstChild.FirstChild.NextSibling
		for depth, n := range xhtml.DescendantsDepth(body) {
			ntype := "element"
			if n.Type != html.ElementNode {
				ntype = "text"
			}
			fmt.Println("depth:", depth, "type:", ntype, "data:", n.Data)
		}
	}
	// Output:
	// DescendantsDepth("<div><span></span><span>a</span></div>")
	// depth: 1 type: element data: div
	// depth: 2 type: element data: span
	// depth: 2 type: element data: span
	// depth: 3 type: text data: a
	// DescendantsDepth("<div><span><span>a</span></span></div>")
	// depth: 1 type: element data: div
	// depth: 2 type: element data: span
	// depth: 3 type: element data: span
	// depth: 4 type: text data: a
}
