package xhtml

import (
	"iter"
	"slices"

	"golang.org/x/net/html"
)

// ShallowEqual returns true
// if a and b have the same Type, DataAtom, Data, Namespace, and Attr.
// It does not look at parents, children, or siblings.
// nil *html.Nodes are never equal to each other.
func ShallowEqual(a, b *html.Node) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Type != b.Type ||
		a.DataAtom != b.DataAtom ||
		a.Namespace != b.Namespace ||
		len(a.Attr) != len(b.Attr) ||
		a.Data != b.Data {
		return false
	}
	if !slices.Equal(a.Attr, b.Attr) {
		return false
	}
	return true
}

// DeepEqual returns true if a and b are [ShallowEqual]
// and all of their descendants are ShallowEqual as well.
func DeepEqual(a, b *html.Node) bool {
	if !ShallowEqual(a, b) {
		return false
	}
	aKids, stop := iter.Pull2(DescendantsDepth(a))
	defer stop()
	bKids, stop := iter.Pull2(DescendantsDepth(b))
	defer stop()

	for {
		depthA, kidA, okA := aKids()
		depthB, kidB, okB := bKids()
		if okA != okB {
			return false
		}
		if depthA != depthB {
			return false
		}
		if !okA {
			return true
		}
		if !ShallowEqual(kidA, kidB) {
			return false
		}
	}
}
