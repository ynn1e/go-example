package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// function has a type parameter(TP)

type Yen int

func (y Yen) String() string {
	return fmt.Sprintf("%d円", y)
}

type OK bool

func (ok OK) String() string {
	if ok {
		return "OK"
	}
	return "NG"
}

func ToStrings[TP fmt.Stringer](v []TP) []string {
	strs := make([]string, len(v))
	for i := range v {
		strs[i] = v[i].String()
	}
	return strs
}

// type has a type parameter(TP)

type List[TP comparable] []TP

func NewList[TP comparable](v ...TP) *List[TP] {
	l := make(List[TP], len(v))
	copy(l, v)
	return &l
}

// method must have no type parameters
/*
func (l List[TP]) GetLast[TP comparable]() (TP, bool) {
	if len(l) == 0 {
		return *new(TP), false
	}
	return l[len(l)-1], true
}
*/

// need to implement like this
func GetLast[TP comparable](l List[TP]) (TP, bool) {
	if len(l) == 0 {
		return *new(TP), false
	}
	return l[len(l)-1], true
}

func (l *List[C]) include(v C) bool {
	for _, lv := range *l {
		if lv == v {
			return true
		}
	}
	return false
}

// unions

func Max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

func main() {
	/*
		$ go run main.go
		[1円 10円 100円]
		[NG]
	*/
	fmt.Println(ToStrings([]Yen{1, 10, 100}))
	fmt.Println(ToStrings([]OK{false}))
	fmt.Println()

	/*
		true
		false
		1 true
	*/
	l := NewList(100, 100, 50, 10, 10, 1, 1, 1, 1) // func NewList(v ...int) *List[int] // func[TP comparable](v ...TP) *List[TP]
	NewList(1.4, 1.2, 3.8)                         // func NewList(v ...float64) *List[float64] // func[TP comparable](v ...TP) *List[TP]
	NewList[bool](false, false, true)              // can specify type(e.g. bool like this) func NewList(v ...bool) *List[bool] // func[TP comparable](v ...TP) *List[TP]
	fmt.Println(l.include(50))
	fmt.Println(l.include(5))
	fmt.Println(GetLast(*l))
	fmt.Println()

	/*
		2
		3.2
	*/
	fmt.Println(Max(1, 2))
	fmt.Println(Max(3.2, 2.8))
}
