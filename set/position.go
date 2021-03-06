package set

import (
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
)

//Layer sets the z-index of this seed.
//this will influence the rendering order.
func Layer(layer int) Style {
	return css.SetZIndex(css.Int(layer))
}

//Side is an sied used to specify how to attach to a parent with Sticky and Overlay.
type side int

//Direction constants.
const (
	Top side = iota
	Bottom
	Left
	Right
)

//Sticky keeps this seed on the screen when scrolling.
func Sticky(attachto ...side) Style {
	if len(attachto) == 0 {
		attachto = []side{Top, Left}
	}

	var rules css.Rules
	for _, side := range attachto {
		switch side {
		case Top:
			rules = append(rules, css.SetTop(css.Zero))
		case Bottom:
			rules = append(rules, css.SetBottom(css.Zero))
		case Left:
			rules = append(rules, css.SetLeft(css.Zero))
		case Right:
			rules = append(rules, css.SetRight(css.Zero))
		}
	}

	return append(rules, css.SetPosition(css.Sticky))
}

//Overlay positions this seed overlayed on its parent.
func Overlay(attachto ...side) Style {
	if len(attachto) == 0 {
		attachto = []side{Top, Left}
	}

	var rules css.Rules
	for _, side := range attachto {
		switch side {
		case Top:
			rules = append(rules, css.SetTop(css.Zero))
		case Bottom:
			rules = append(rules, css.SetBottom(css.Zero))
		case Left:
			rules = append(rules, css.SetLeft(css.Zero))
		case Right:
			rules = append(rules, css.SetRight(css.Zero))
		}
	}

	return append(rules, css.SetPosition(css.Absolute))
}

//OffsetTop sets the top offset of this seed.
func OffsetTop(u units.Unit) Style {
	return css.SetTop(css.Measure(u))
}

//OffsetBottom sets the bottom offset of this seed.
func OffsetBottom(u units.Unit) Style {
	return css.SetBottom(css.Measure(u))
}

//OffsetLeft sets the offset to the left of this seed.
func OffsetLeft(u units.Unit) Style {
	return css.SetLeft(css.Measure(u))
}

//OffsetRight sets the offset to the right of this seed.
func OffsetRight(u units.Unit) Style {
	return css.SetRight(css.Measure(u))
}
