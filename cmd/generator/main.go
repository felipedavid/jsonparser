package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ./cmd <seed> <n_points>\n")
		return
	}

	seed, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic("Unable to parse seeds")
	}

	r := rand.New(rand.NewSource(seed))

	nPoints, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Unable to parse points")
	}

	fmt.Println(`{"pairs":[`)
	for i := 0; i < nPoints; i++ {
		if i != 0 {
			fmt.Printf(",\n")
		}
		x0 := (r.Float64() * 360) - 180
		y0 := (r.Float64() * 180) - 90
		x1 := (r.Float64() * 360) - 180
		y1 := (r.Float64() * 180) - 90
		fmt.Print("\t")
		fmt.Printf(`{"x0":%v, "y0":%v, "x1":%v, "y1":%v}`, x0, y0, x1, y1)
	}
	fmt.Print("\n]}\n")
}
