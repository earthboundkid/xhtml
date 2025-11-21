package xhtml

import (
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// AbsolutizeURLs mutates the document
// so that any href or src attributes on appropriate elements
// are absolute URLs as resolved by base.
//
// Doesn't parse JavaScript or CSS.
func AbsolutizeURLs(n *html.Node, base *url.URL) {
	for c := range n.Descendants() {
		switch c.DataAtom {
		case atom.A, atom.Link, atom.Area:
			if href := Attr(c, "href"); href != "" {
				if attrURL, err := base.Parse(href); err == nil {
					SetAttr(c, "href", attrURL.String())
				}
			}
		case atom.Img, atom.Script, atom.Iframe, atom.Audio, atom.Video, atom.Source, atom.Embed:
			if src := Attr(c, "src"); src != "" {
				if attrURL, err := base.Parse(src); err == nil {
					SetAttr(c, "src", attrURL.String())
				}
			}
		}
	}
}
