package main

import "fmt"

type Fruit int

const (
	Unknown Fruit = iota
	Apples
	Bananas
	Cherries
	Durian
	Elderberry
)

// use stringer: https://pkg.go.dev/fmt#Stringer
func (f Fruit) String() string {
	switch f {
	case Apples:
		return "Apples"
	case Bananas:
		return "Bananas"
	case Cherries:
		return "Cherries"
	case Durian:
		return "Durian"
	case Elderberry:
		return "Elderberry"
	default:
		return "Unknown"
	}
}

/*
Apples
Durian
2
100
Unknown
*/
func main() {
	fmt.Println(Apples)
	fmt.Println(Durian)
	fmt.Println(2)
	fmt.Println(100)

	var kiwi Fruit = 11

	fmt.Println(kiwi)
}