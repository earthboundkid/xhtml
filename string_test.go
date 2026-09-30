package xhtml_test

import (
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/xhtml"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestInnerText(t *testing.T) {
	be := assert.FailsNow(t)
	for _, tc := range []struct {
		input, want string
	}{
		{
			`<p><em></em>&lt;div class=&#34;flourish-embed flourish-cards&#34; data-src=&#34;visualisation/14836391&#34;&gt;&lt;script src=&#34;<a href="https://public.flourish.studio/resources/embed.js">https://public.flourish.studio/resources/embed.js</a>&#34;&gt;&lt;/script&gt;&lt;/div&gt;
	</p>`,
			`<div class="flourish-embed flourish-cards" data-src="visualisation/14836391"><script src="https://public.flourish.studio/resources/embed.js"></script></div>`,
		},
	} {
		doc := be.OK(html.Parse(strings.NewReader(tc.input)))
		p := xhtml.Select(doc, xhtml.WithAtom(atom.P))
		got := xhtml.TextContent(p)
		be.Equal(got, tc.want)
	}
}
