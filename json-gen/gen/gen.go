package gen

//go:generate go run golang.org/x/tools/cmd/stringer -type=Meat
type Meat int

const (
	Beef Meat = iota
	Pork
	Chicken
	Mutton
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Recipe
type Recipe int

const (
	Grill Recipe = iota
	Roast
	Broil
	Fly
	Steam
	Sear
)