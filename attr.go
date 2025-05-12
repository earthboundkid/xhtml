package xhtml

import (
	"slices"

	"golang.org/x/net/html"
)

// Attr returns the value of the first attribute of n with the named key.
func Attr(n *html.Node, name string) string {
	if n == nil {
		return ""
	}
	for _, attr := range n.Attr {
		if attr.Key == name && attr.Namespace == "" {
			return attr.Val
		}
	}
	return ""
}

// SetAttr adds or replaces the give attribute key and value on n.
func SetAttr(n *html.Node, key, value string) {
	for i := range n.Attr {
		attr := &n.Attr[i]
		if attr.Key == key && attr.Namespace == "" {
			attr.Val = value
			return
		}
	}
	n.Attr = append(n.Attr, html.Attribute{
		Key: key,
		Val: value,
	})
}

// DeleteAttr removes any attributes of n with the named key.
func DeleteAttr(n *html.Node, key string) {
	n.Attr = slices.DeleteFunc(n.Attr, func(a html.Attribute) bool {
		return a.Key == key && a.Namespace == ""
	})
}
