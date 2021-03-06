package markdown

import (
	"regexp"

	"qlova.org/seed"
	"qlova.org/seed/use/html"
	"qlova.org/seed/new/html/div"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.UGCPolicy()

func init() {
	policy.AllowAttrs("style").OnElements("span", "p")
	policy.AllowStyles("color").Matching(regexp.MustCompile("(?i)^#([0-9a-f]{3,4}|[0-9a-f]{6}|[0-9a-f]{8})$")).Globally()
}

//New returns a new markdown container.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		seed.Options(options),
	)
}

//Set sets the inner HTML of the seed to rendered and sanitized markdown.
func Set(md string) seed.Option {
	rendered := markdown.ToHTML([]byte(md), nil, nil)
	rendered = policy.SanitizeBytes(rendered)
	return html.Set(string(rendered))
}
