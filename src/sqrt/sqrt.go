package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(Sqrt(f))
}

// Sqrt ... calculates te approximate square root
func Sqrt(n float64) float64 {
	var z, zx, err, done = 1.0, 0.0, 0.0000000000000005, false
	for !done {
		z -= (z*z - n) / (2 * z)
		// fmt.Printf("z is %g, zx = %g\n", z, zx)
		if math.Abs(z-zx) < err {
			done = true
		}
		zx = z
	}
	return z
}
