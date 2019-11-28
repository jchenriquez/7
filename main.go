package main

import (
	"fmt"
	"math"
)

var UPPER_BOUND = int(math.Pow(2, 32))-1
var LOWER_BOUND = -int(math.Pow(2, 32))

func reverse(x int) int {
	var res int
	if x == 0 {
		return 0
	}
	sign := x / int(math.Abs(float64(x)))
	x = int(math.Abs(float64(x)))
	tens := int(math.Pow(10, math.Floor(math.Log10(float64(x)))))

	for x > 0 {
		rem := int(math.Mod(float64(x), 10))
		res += rem*tens
		x /= 10
		tens/=10
	}

	res *= sign

	if res < LOWER_BOUND || res > UPPER_BOUND {
		return 0
	}

	return res
}

func main() {
	fmt.Printf("result %d\n", reverse(1534236469))
}
