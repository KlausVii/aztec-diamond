package aztec

import (
	"image/color"
	"strings"

	"github.com/fogleman/gg"
)

func (d *Diamond) String() string {
	le := len(d.g)

	b := strings.Builder{}

	c := 0
	for i, l := range d.g {
		if i != c {
			b.WriteString("\n")
			c = i
		}
		for j, p := range l {
			if j == 0 {
				b.WriteString(strings.Repeat(" ", (le-len(l))/2))
			}

			switch p {
			case none:
				b.WriteString("*")
			case up:
				b.WriteString("^")
			case down:
				b.WriteString("v")
			case left:
				b.WriteString("<")
			case right:
				b.WriteString(">")
			}
		}
	}

	return b.String()
}

func (d *Diamond) Draw() *gg.Context {
	le := len(d.g)

	dc := gg.NewContext(20*le, 20*le)

	for i, l := range d.g {
		offset := float64(20 * (le - len(l)) / 2)
		for j, p := range l {
			dc.DrawRectangle(offset+float64(j)*20, float64(i)*20, 20, 20)
			switch p {
			case none:
				dc.SetColor(color.White)
			case up:
				dc.SetRGB255(255, 0, 0)
			case down:
				dc.SetRGB255(0, 255, 0)
			case left:
				dc.SetRGB255(0, 0, 255)
			case right:
				dc.SetRGB255(255, 255, 0)
			}
			dc.Fill()
		}
	}

	return dc
}
