package aztec

import (
	"log"
	"math/rand"
)

type direction int

const (
	none direction = iota
	up
	down
	left
	right
)

type Diamond struct {
	g [][]direction
}

func NewDiamond() *Diamond {
	d := &Diamond{g: newGrid(1)}

	return d.Fill()
}

func newGrid(n int) [][]direction {
	if n == 0 {
		return nil
	}
	g := make([][]direction, 0, n*2)

	for i := 1; i <= n; i++ {
		g = append(g, make([]direction, i*2))
	}

	for i := n; i > 0; i-- {
		g = append(g, make([]direction, i*2))
	}

	return g
}

func (d *Diamond) Iter(n int) *Diamond {
	for i := 0; i < n; i++ {
		d.Grow().Fill()
	}
	return d
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
			if p == none {
				// vertical dominoes
				if rand.Int31n(2) == 1 {
					first := left
					second := right

					d.g[i][j] = first
					d.g[i+1][j+offset] = first
					d.g[i][j+1] = second
					d.g[i+1][j+offset+1] = second
					continue
				}

				// horizontal dominoes
				first := up
				second := down
				d.g[i][j] = first
				d.g[i][j+1] = first
				d.g[i+1][j+offset] = second
				d.g[i+1][j+offset+1] = second
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
		switch i {
		case offsetSwitch:
			downOff = 0
		case offsetSwitch + 1:
			downOff = -1
			upOff = 0
		case offsetSwitch + 2:
			upOff = -1
		}

		rowLen := len(row)
		for j := 0; j < rowLen; j++ {
			p := d.g[i][j]
			var x, y int
			switch p {
			case up:
				if (upOff <= 0 || (j != 0 && j != rowLen-1)) && d.g[i-1][j-upOff] == down {
					//skip next point too
					j++
					continue
				}
				x, y = i, j+1-upOff

			case down:
				if (downOff >= 0 || (j != 0 && j != rowLen-1)) && d.g[i+1][j+downOff] == up {
					//skip next point too
					j++
					continue
				}
				x, y = i+2, j+1+downOff
			case left:
				if j != 0 && d.g[i][j-1] == right {
					continue
				}
				x, y = i+1, j

			case right:
				if j != rowLen-1 && d.g[i][j+1] == left {
					continue
				}
				x, y = i+1, j+2
			}

			if next[x][y] != none {
				log.Panicf("%d, %d not nil", x, y)
			}
			next[x][y] = p
		}
	}

	d.g = next

	return d
}
