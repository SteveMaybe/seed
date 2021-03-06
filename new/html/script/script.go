package script

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("script"), seed.Options(options))
}
