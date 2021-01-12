package aztec

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type domino struct {
	vertical bool
	direction
}

type Diamond struct {
	g [][]*domino
}

func NewDiamond(n int) *Diamond {
	return &Diamond{g: newGrid(n)}
}

func newGrid(n int) [][]*domino {
	if n == 0 {
		return nil
	}
	g := make([][]*domino, 0, n*2)

	for i := 1; i <= n; i++ {
		g = append(g, make([]*domino, i*2))
	}

	for i := n; i > 0; i-- {
		g = append(g, make([]*domino, i*2))
	}

	return g
}

func (d *Diamond) Iter(n int, draw bool) *Diamond {
	for i := 0; i < n; i++ {
		d.Grow()
		if draw {
			fmt.Printf("%s\n\n", d.Draw())
		}
		d.Fill()
		if draw {
			fmt.Printf("%s\n\n", d.Draw())
		}
	}
	return d
}

func (d *Diamond) Draw() string {
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
			if p == nil {
				b.WriteString("*")
				continue
			}
			switch p.direction {
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

func (d *Diamond) Fill() *Diamond {
	offset := 1
	offsetSwitch := len(d.g)/2 - 1
	for i := 0; i < len(d.g); i++ {
		if i == offsetSwitch {
			offset = 0
		}
		if i == offsetSwitch+1 {
			offset = -1
		}

		for j := 0; j < len(d.g[i]); j++ {
			p := d.g[i][j]
			if p == nil {
				vert := rand.Int31n(2) == 1
				first := newDomino(vert)
				second := newDomino(vert)
				d.g[i][j] = first
				if vert {
					first.direction = left
					second.direction = right
					d.g[i+1][j+offset] = first
					d.g[i][j+1] = second
					d.g[i+1][j+offset+1] = second

				} else {
					first.direction = up
					second.direction = down
					d.g[i][j+1] = first
					d.g[i+1][j+offset] = second
					d.g[i+1][j+offset+1] = second
				}
			}
		}
	}

	return d
}

func (d *Diamond) Grow() *Diamond {
	next := newGrid(len(d.g)/2 + 1)
	downOff := 1
	upOff := 1
	offsetSwitch := len(d.g)/2 - 1

	for i, row := range d.g {
		if i == offsetSwitch {
			downOff = 0
		}
		if i == offsetSwitch+1 {
			downOff = -1
			upOff = 0
		}
		if i == offsetSwitch+2 {
			upOff = -1
		}
		rowLen := len(row)
		for j := 0; j < rowLen; j++ {
			p := d.g[i][j]
			var x, y int
			switch p.direction {
			case up:
				if (upOff <= 0 || (j != 0 && j != rowLen-1)) && d.g[i-1][j-upOff].direction == down {
					//skip next point too
					j++
					continue
				}
				x, y = i, j+1-upOff

			case down:
				if (downOff >= 0 || (j != 0 && j != rowLen-1)) && d.g[i+1][j+downOff].direction == up {
					//skip next point too
					j++
					continue
				}
				x, y = i+2, j+1+downOff
			case left:
				if j != 0 && d.g[i][j-1].direction == right {
					continue
				}
				x, y = i+1, j

			case right:
				if j != rowLen-1 && d.g[i][j+1].direction == left {
					continue
				}
				x, y = i+1, j+2
			}

			if next[x][y] != nil {
				log.Panicf("%d, %d not nil", x, y)
			}
			next[x][y] = p
		}
	}

	d.g = next

	return d
}

func newDomino(vert bool) *domino {
	return &domino{
		vertical: vert,
	}
}
