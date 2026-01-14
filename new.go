package xhtml

import (
	"slices"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// New creates a new [html.Node] with the specified tag and attributes.
// It handles properly setting the Node.Type and Node.DataAtom.
// New panics if the length of attrs is not even.
func New(tag string, attrs ...string) *html.Node {
	var attrslice []html.Attribute
	if len(attrs) > 0 {
		if len(attrs)%2 != 0 {
			panic("uneven number of attr/value pairs")
		}
		attrslice = make([]html.Attribute, len(attrs)/2)
		for i := range attrslice {
			attrslice[i] = html.Attribute{
				Key: attrs[i*2],
				Val: attrs[i*2+1],
			}
		}
	}
	dataAtom := atom.Lookup([]byte(tag))
	// Consolidate memory usage if possible
	if dataAtom != 0 {
		tag = dataAtom.String()
	}
	return &html.Node{
		Type:     html.ElementNode,
		Data:     tag,
		DataAtom: dataAtom,
		Attr:     attrslice,
	}
}

// LastChildOrNew returns the last child of p
// if it is [ShallowEqual] to a new *html.Node with tag and attrs.
// Otherwise, it appends a new *html.Node with tag and attrs
// and returns that.
// For why this is operation useful,
// see [Converting docx to clean HTML].
//
// [Converting docx to clean HTML]: https://mike.zwobble.org/2013/12/docx-to-clean-html-handling-the-xml-structure-mismatch/
func LastChildOrNew(p *html.Node, tag string, attrs ...string) *html.Node {
	n := New(tag, attrs...)
	if ShallowEqual(p.LastChild, n) {
		return p.LastChild
	}
	p.AppendChild(n)
	return n
}

// Clone n and all of its children.
func Clone(n *html.Node) *html.Node {
	new := &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      slices.Clone(n.Attr),
	}
	for c := range n.ChildNodes() {
		c2 := Clone(c)
		new.AppendChild(c2)
	}
	return new
}
