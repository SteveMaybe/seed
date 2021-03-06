//Package dropdown provides an input widget for selecting from an array of preset values.
package dropdown

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/html/htmlselect"
	"qlova.org/seed/new/html/option"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/text/rich"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"
)

type data struct {
	values []string

	feed *feed.Feed

	placeholder string
}

//New returns a new dropdown.
func New(options ...seed.Option) seed.Seed {

	var data data

	var c = htmlselect.New(options...)

	c.Load(&data)

	if data.placeholder != "" {
		option.New(
			text.Set(rich.Text(data.placeholder)),
			attr.Set("disabled", ""),
			attr.Set("selected", ""),
			attr.Set("hidden", ""),
			attr.Set("value", ""),
		).AddTo(c)
	}

	if data.feed != nil {
		data.feed.Mutate(html.SetTag("select"), seed.Options(options))
		c = data.feed.New(
			option.New(text.SetStringTo(js.String{Value: data.feed.Data.Value})),
		)
	}

	for _, val := range data.values {
		option.New(text.Set(rich.Text(val))).AddTo(c)
	}

	return c
}

//Update updates the given variable whenever the dropdown value is modified.
func Update(variable *clientside.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("change", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
			client.On("input", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
		)
	})
}

//Set sets the preset dropdown values.
func Set(values []string) seed.Option {
	return seed.Mutate(func(d *data) {
		d.values = values
	})
}

//Feed feeds the dropdown with the given feed.
func Feed(f *feed.Feed) seed.Option {
	return seed.Mutate(func(d *data) {
		d.feed = f
	})
}

//SetPlaceholder sets the placeholder text of the dropbox.
func SetPlaceholder(placeholder string) seed.Option {
	return seed.Mutate(func(d *data) {
		d.placeholder = placeholder
	})
}
