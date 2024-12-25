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
