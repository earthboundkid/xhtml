// Package xhtml makes x/net/html easier
package xhtml

import (
	"iter"
	"regexp"
	"slices"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Selector is a function that matches html.Nodes.
type Selector func(n *html.Node) bool

// And composes two selectors together with &&.
func (s Selector) And(other Selector) Selector {
	return func(n *html.Node) bool {
		return s(n) && other(n)
	}
}

// Or composes two selectors together with ||.
func (s Selector) Or(other Selector) Selector {
	return func(n *html.Node) bool {
		return s(n) || other(n)
	}
}

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

// WithAttr returns a Selector that matches html.Nodes with the given attribute key and value.
func WithAttr(key, value string) Selector {
	return func(n *html.Node) bool {
		return Attr(n, key) == value
	}
}

// WithAttrKey returns a Selector that matches html.Nodes with the given attribute key.
// It matches even if the value is empty.
func WithAttrKey(attr string) Selector {
	return func(n *html.Node) bool {
		return slices.ContainsFunc(n.Attr, func(a html.Attribute) bool {
			return a.Key == attr
		})
	}
}

// WithAttrMatch returns a Selector
// that matches html.Nodes where the given attribute key's value
// matches the provided regex pattern.
//
// WithAttrMatch panics if the pattern does not compile.
func WithAttrMatch(key, pattern string) Selector {
	re := regexp.MustCompile(pattern)
	return func(n *html.Node) bool {
		return re.MatchString(Attr(n, key))
	}
}

// WithID returns a Selector that matches html.Nodes with the given id= attribute.
func WithID(id string) Selector {
	return WithAttr("id", id)
}

// WithDataset returns a Selector that matches html.Nodes with the given data attribute set.
// The attribute should be in kebab-case, not camelCase, without the "data-" prefix.
func WithDataset(attr string) Selector {
	return WithAttrKey("data-" + attr)
}

// DescendantsDepth is like html.Node.Descendants(),
// except it also yields the depth of each node relative to its parent.
func DescendantsDepth(n *html.Node) iter.Seq2[int, *html.Node] {
	_ = n.FirstChild // eager nil check

	return func(yield func(int, *html.Node) bool) {
		descendants(1, n, yield)
	}
}

func descendants(depth int, n *html.Node, yield func(int, *html.Node) bool) bool {
	for c := range n.ChildNodes() {
		if !yield(depth, c) || !descendants(depth+1, c, yield) {
			return false
		}
	}
	return true
}
