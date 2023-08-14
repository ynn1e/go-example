package stat

import "math"

func PoissonDistribution(avg, n int) float64 {
	return math.Pow(float64(avg), float64(n)) * math.Pow(CreateNapiersConstant(50), -float64(avg)) / float64(Factorial(n))
}

func CreateNapiersConstant(n int) float64 {
	e := 2.0

	for i := 2; i <= n; i++ {
		e += 1.0 / float64(Factorial(i))
	}

	return e
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return Factorial(n-1) * n
}
