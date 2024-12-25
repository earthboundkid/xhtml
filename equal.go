package xhtml

import "golang.org/x/net/html"

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
	for i := range a.Attr {
		if a.Attr[i] != b.Attr[i] {
			return false
		}
	}
	return true
}
