//go:build go1.24

package xhtml

import (
	"strings"

	"golang.org/x/net/html"
)

// WithClass returns a Selector that matches nodes with classname.
func WithClass(classname string) Selector {
	return func(n *html.Node) bool {
		for field := range strings.FieldsSeq(Attr(n, "class")) {
			if field == classname {
				return true
			}
		}
		return false
	}
}
