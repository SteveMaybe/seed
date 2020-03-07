package anime

import (
	"strconv"
	"time"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

var id int

//Animation can be added to a seed as the current playing animation of that seed.
type Animation struct {
	id int

	keyframes Keyframes
	options   []seed.Option
}

func (anim Animation) AddTo(c seed.Any) {
	data := seeds[c.Root()]
	data.animations = append(data.animations, anim)
	seeds[c.Root()] = data

	css.SetAnimationName(css.AnimationName("a" + strconv.Itoa(anim.id))).AddTo(c)
	for _, o := range anim.options {
		o.AddTo(c)
	}
}

func (anim Animation) Apply(c seed.Ctx) {
	data := seeds[c.Root()]
	data.animations = append(data.animations, anim)
	seeds[c.Root()] = data

	css.SetAnimationName(css.AnimationName("a" + strconv.Itoa(anim.id))).Apply(c)
	for _, o := range anim.options {
		o.Apply(c)
	}
}

func (anim Animation) Reset(c seed.Ctx) {
	data := seeds[c.Root()]
	data.animations = append(data.animations, anim)
	seeds[c.Root()] = data

	css.SetAnimationName(css.AnimationName("a" + strconv.Itoa(anim.id))).Reset(c)
	for _, o := range anim.options {
		o.Reset(c)
	}
}

func (anim Animation) And(more ...seed.Option) seed.Option {
	return seed.And(anim, more...)
}

type data struct {
	animations []Animation
}

var seeds = make(map[seed.Seed]data)

func Set(styles ...css.Style) KeyFrame {
	return KeyFrame(styles)
}

type KeyFrame []css.Style

func (k KeyFrame) Rules() css.Rules {
	var result css.Rules
	for _, s := range k {
		result = append(result, s.Rules()...)
	}
	return result
}

type Keyframes map[float32]css.Style

//New returns a new animation from the given keyframes.
//It also takes a variable amount of options that will be applied to any seed with this animation set.
//These are useful for setting default animation durations, looping etc.
func New(keyframes Keyframes, options ...seed.Option) Animation {
	id++
	return Animation{
		id,
		keyframes,
		options,
	}
}

//Loop sets the current playing animation to loop back and forth.
func Loop() css.Rules {
	return css.Rules{
		css.SetAnimationDirection(css.Alternate),
		css.SetAnimationIterationCount(css.Infinite),
	}
}

//SetDuration sets the duration of the current playing animation.
func SetDuration(d time.Duration) css.Rule {
	return css.SetAnimationDuration(css.Duration(d))
}
