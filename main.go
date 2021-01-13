package main // import "github.com/KlausVii/aztec-diamond"

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/KlausVii/aztec-diamond/aztec"
)

var out = flag.String("out", "out.png", "path to output")
var order = flag.Int("order", 1, "order of the tiling")
var seed = flag.Int64("seed", time.Now().UnixNano(), "optional seed, defaults to unix nano")

func main() {
	flag.Parse()
	rand.Seed(*seed)

	if *order < 1 {
		log.Fatalf("order must be greater than or equal to 1")
	}

	d := aztec.NewDiamond().Iter(int(*order - 1))

	dc := d.Draw()
	if err := dc.SavePNG(*out); err != nil {
		log.Fatalf("failed to save output to %s", *out)
	}
}
