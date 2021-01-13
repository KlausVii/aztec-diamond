package aztec

import (
	"math/rand"
	"testing"
)

func TestDiamond_Iter(t *testing.T) {
	expected := `          ^^
         ^^^^
        ^^^^^^
       ^^<>^^^^
      ^^<<>>^^^^
     <><<^^>><><>
    <<><^^^^><><>>
   <<^^<><vv>^^^^>>
  <<<vv<><^^>>^^^^>>
 <<<<<><>^^^^>vv><>>>
<<<<<<><>>^^^^^^><>>>>
<<<<<^^<>>vvvv><vv>>>>
 <<<^^<<>vv>^^><^^>>>
  <<vv<^^^^>vvvvvv>>
   <^^<vv><vvvv>^^>
    vv<^^><<>^^>vv
     vvvvvv<>vvvv
      vvvvvvvvvv
       vvvvvvvv
        vvvvvv
         vvvv
          vv`

	rand.Seed(1123142)
	d := NewDiamond().Fill().Iter(10)

	if d.String() != expected {
		t.Fail()
	}
}
