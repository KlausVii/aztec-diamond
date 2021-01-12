package main // import "github.com/KlausVii/aztec-diamond"

import (
	"fmt"
	"math/rand"

	"github.com/KlausVii/aztec-diamond/aztec"
)

func main() {
	rand.Seed(1123142)
	d := aztec.NewDiamond(1).Fill().Iter(128, false)

	fmt.Println(d.Draw())
}
