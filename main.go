package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int
	tens := int(math.Pow(10, math.Floor(math.Log10(float64(x)))))

	for x > 0 {
		rem := int(math.Mod(float64(x), 10))
		fmt.Printf("tens %d\n", tens)
		res += rem*tens
		x /= 10
		tens/=10
	}

	return res
}

func main() {
	fmt.Printf("result %d\n", reverse(1269))
}
