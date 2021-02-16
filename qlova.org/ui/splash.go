package ui

import (
	"math"
	"time"

	"qlova.tech/rgb"

	"qlova.org/seed"
	"qlova.org/seed/new/animation"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/markdown"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"

	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/transition"

	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units/rem"
	"qlova.org/seed/use/css/units/vh"
)

//QlovaSplash splashes the Qlova logo.
type QlovaSplash struct{}

//Page implements page.Page
func (QlovaSplash) Page(r page.Router) seed.Seed {
	return page.New(
		transition.Fade(),

		spacer.New(rem.New(2)),

		//Qlova and spinning logo.
		row.New(align.Center(), set.Width(rem.New(22)),

			set.OffsetLeft(rem.New(-2)),

			animation.New(
				animation.Duration(time.Second*2),
				animation.Frames{
					0:   set.Translation(nil, vh.New(30)),
					40:  set.Translation(nil, vh.New(30)),
					100: set.Translation(nil, nil),
				},
			),
			css.SetAnimationFillMode(css.Forwards),
			css.Set("animation-timing-function", "ease-in-out"),

			image.New(set.Height(rem.New(6)),
				image.Set("logo.svg"),

				animation.New(
					animation.Duration(time.Second*2),
					animation.Frames{
						0:   set.Angle(-math.Pi * 4),
						40:  set.Angle(-math.Pi * 4),
						100: set.Angle(0),
					},
				),

				css.Set("animation-timing-function", "ease-in-out"),
				css.SetAnimationDirection(css.Normal),
			),
			spacer.New(rem.New(0.5)),
			text.New(text.SetColor(rgb.Hex("#565656")), text.SetSize(rem.New(6)),
				text.Set("Qlova"),
			),
			spacer.New(rem.New(5)),
		),

		spacer.New(rem.New(2)),

		//Description.
		markdown.New(align.Center(), set.MaxWidth(rem.New(35)),

			set.Opacity(0),
			animation.New(
				animation.Duration(time.Second),
				animation.Frames{
					0:   animation.Frame{set.Translation(nil, vh.New(25)), set.Opacity(0)},
					100: animation.Frame{set.Translation(nil, nil), set.Opacity(1)},
				},
			),
			css.SetAnimationFillMode(css.Forwards),
			css.Set("animation-delay", "1s"),
			css.Set("animation-timing-function", "ease-in-out"),

			text.SetSize(rem.New(1.5)),
			text.SetColor(rgb.Hex("#565656")),

			markdown.Set(`
Welcome! your probably looking for [Qlovaseed](https://qlovaseed.com).
That's our free app development module for Go.

We also have more projects coming down the pipeline, 
so do check out our [Github](https://github.com) page.

You can also find us on [Twitter](https://twitter.com/QlovaNZ) & [Facebook](https://facebook.com/QlovaNZ).
			`)),
	)
}
