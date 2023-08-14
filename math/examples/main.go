package main

import (
	"fmt"

	"example.com/stat"
)

func main() {
	fmt.Printf("5!=%d\n", stat.Factorial(5))
	fmt.Printf("-1!=%d\n", stat.Factorial(-1))

	fmt.Printf("ex=%f\n", stat.CreateNapiersConstant(50))

	for i := 0; i < 10; i++ {
		fmt.Printf("poisson distribution: avg=3, n=%d, val=%f\n", i, stat.PoissonDistribution(3, i))
	}
}
