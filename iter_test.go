package xhtml_test

import (
	"fmt"
	"strings"

	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ExampleSelectAll() {
	{
		doc, err := html.Parse(strings.NewReader(`
			<ul>
				<li><a href="https://example.com/en-us">Hello, World!</a></li>
				<li><a href="https://example.com/ja-jp">こんにちは世界！</a></li>
			</ul>`))
		if err != nil {
			panic(err)
		}
		// Find all a[href] in a document
		// And print the link URL and text
		for link := range xhtml.SelectAll(doc,
			xhtml.WithAtom(atom.A).And(xhtml.WithAttrKey("href"))) {
			fmt.Println(xhtml.Attr(link, "href"), xhtml.TextContent(link))
		}
	}
	{
		doc, err := html.Parse(strings.NewReader(`
			<ul>
				<li><a href="https://example.com/a/something">Link A</a></li>
				<li><a href="https://example.com/b/something-else">Link B</a></li>
			</ul>
		`))
		if err != nil {
			panic(err)
		}

		// Find all a[href^=https://example.com/a] in a document
		// And print the link URL and text
		for link := range xhtml.SelectAll(doc,
			xhtml.WithAtom(atom.A).And(xhtml.WithAttrMatch("href", `^https://example.com/a`))) {
			fmt.Println(xhtml.Attr(link, "href"), xhtml.TextContent(link))
		}

	}
	{
		doc, err := html.Parse(strings.NewReader(`
			<em class="italic text-xl">Sign up</em>
			<strong class="bold text-2xl">today</strong>
			<span>for up to</span>
			<strong class="text-xl bold">50% off</strong>`))
		if err != nil {
			panic(err)
		}
		// Find all em or strong with class text-xl in a document
		// And print the text
		for tag := range xhtml.SelectAll(doc,
			xhtml.WithAtom(atom.Em).Or(xhtml.WithAtom(atom.Strong)).And(xhtml.WithClass("text-xl"))) {
			fmt.Println(xhtml.TextContent(tag))
		}
	}
	// Output:
	// https://example.com/en-us Hello, World!
	// https://example.com/ja-jp こんにちは世界！
	// https://example.com/a/something Link A
	// Sign up
	// 50% off
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
			fmt.Println("depth:", depth, "type:", n.Type, "data:", n.Data)
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
			fmt.Println("depth:", depth, "type:", n.Type, "data:", n.Data)
		}
	}
	// Output:
	// DescendantsDepth("<div><span></span><span>a</span></div>")
	// depth: 1 type: ElementNode data: div
	// depth: 2 type: ElementNode data: span
	// depth: 2 type: ElementNode data: span
	// depth: 3 type: TextNode data: a
	// DescendantsDepth("<div><span><span>a</span></span></div>")
	// depth: 1 type: ElementNode data: div
	// depth: 2 type: ElementNode data: span
	// depth: 3 type: ElementNode data: span
	// depth: 4 type: TextNode data: a
}
