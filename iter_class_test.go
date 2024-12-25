//go:build go1.24

package xhtml_test

import (
	"fmt"
	"strings"

	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
)

func ExampleWithClass() {
	doc, err := html.Parse(strings.NewReader(`
	<ul>
		<li><a class="bttn english" href="https://example.com/en-us">Hello, World!</a></li>
		<li><a class="bttn japanese" href="https://example.com/ja-jp">こんにちは世界！</a></li>
	</ul>`))
	if err != nil {
		panic(err)
	}
	// Find .english elements
	// And print the link URL and text
	for link := range xhtml.SelectAll(doc, xhtml.WithClass("english")) {
		fmt.Println(xhtml.Attr(link, "href"), xhtml.TextContent(link))
	}
	// Output:
	// https://example.com/en-us Hello, World!
}
