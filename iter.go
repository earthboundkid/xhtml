// Package xhtml makes x/net/html easier
package xhtml

import (
	"iter"
	"slices"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// SelectAll returns an iterator yielding matching nodes.
func SelectAll(n *html.Node, match func(*html.Node) bool) iter.Seq[*html.Node] {
	return func(yield func(*html.Node) bool) {
		for c := range n.Descendants() {
			if match(c) && !yield(c) {
				return
			}
		}
	}
}

// SelectSlice returns a slice of child nodes matched by the selector.
func SelectSlice(n *html.Node, match func(*html.Node) bool) []*html.Node {
	return slices.Collect(SelectAll(n, match))
}

// Select returns the first child node matched by the selector or nil.
func Select(n *html.Node, match func(*html.Node) bool) *html.Node {
	for n := range SelectAll(n, match) {
		return n
	}
	return nil
}

// Closest traverses the node and its parents until it finds a node that matches.
func Closest(n *html.Node, match func(*html.Node) bool) *html.Node {
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

func WithAtom(a atom.Atom) func(n *html.Node) bool {
	return func(n *html.Node) bool {
		return n.DataAtom == a
	}
}

var WithBody = WithAtom(atom.Body)
