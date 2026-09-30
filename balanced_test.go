package xhtml_test

import (
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
)

func TestIsBalanced(t *testing.T) {
	tcs := []struct {
		string
		bool
	}{
		{"", true},
		{"<a></a>", true},
		{"<a><b>hi</b></a>", true},
		{"hello <br /> world", true},
		{"<a><b></b><c></c></a>", true},
		{"</a>", false},
		{"<a></b>", false},
		{"<a><b><c></b></c></a>", false},
	}
	for _, testcase := range tcs {
		assert.FailsNow(t).Run(testcase.string, func(be assert.TB) {
			be.Equal(xhtml.IsBalanced(testcase.string), testcase.bool)
		})
	}
}
