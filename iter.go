// Package xhtml makes x/net/html easier
package xhtml

import (
	"iter"
	"slices"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Selector is a function that matches html.Nodes.
type Selector = func(n *html.Node) bool

// SelectAll returns an iterator yielding matching nodes in n.Descendants().
func SelectAll(n *html.Node, match Selector) iter.Seq[*html.Node] {
	return func(yield func(*html.Node) bool) {
		for c := range n.Descendants() {
			if match(c) && !yield(c) {
				return
			}
		}
	}
}

// SelectSlice returns a slice of descendant nodes matched by the Selector.
func SelectSlice(n *html.Node, match Selector) []*html.Node {
	return slices.Collect(SelectAll(n, match))
}

// Select returns the first descendant node matched by the Selector or nil.
func Select(n *html.Node, match Selector) *html.Node {
	for n := range SelectAll(n, match) {
		return n
	}
	return nil
}

// Closest traverses the node and its parents until it finds a node that matches the Selector.
func Closest(n *html.Node, match Selector) *html.Node {
	if match(n) {
		return n
	}
	for p := range n.Ancestors() {
		if match(p) {
			return p
		}
	}
	return nil
}

// WithAtom returns a Selector that matches html.Nodes with the given [atom.Atom].
func WithAtom(a atom.Atom) Selector {
	return func(n *html.Node) bool {
		return n.DataAtom == a
	}
}

// WithID returns a Selector that matches html.Nodes with the given id= attribute.
func WithID(id string) Selector {
	return func(n *html.Node) bool {
		return Attr(n, "id") == id
	}
}

// WithDataset returns a Selector that matches html.Nodes with the given data attribute set.
// The attribute should be in kebab-case, not camelCase, without the "data-" prefix.
func WithDataset(attr string) Selector {
	attr = "data-" + attr
	return func(n *html.Node) bool {
		return Attr(n, attr) != ""
	}
}
