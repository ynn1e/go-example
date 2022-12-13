package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	jg "github.com/ynn1e/go-example/json-gen/gen"
)

/*
$ go run main.go
{"Meat":0,"Recipe":1}
{Meat:1 Recipe:4}

$ go get golang.org/x/tools/cmd/stringer; go generate; cd gen; go generate; cd ..; go run main.go
{"Meat":0,"Recipe":1}
{Meat:Pork Recipe:Steam}
*/

func main() {
	cooking := struct{
		Meat jg.Meat
		Recipe jg.Recipe
	}{
		jg.Beef, jg.Roast,
	}

	if err := json.NewEncoder(os.Stdout).Encode(cooking); err != nil {
		log.Fatal(err)
	}

	input := `{"Meat":1, "Recipe":4}`
	if err := json.NewDecoder(strings.NewReader(input)).Decode(&cooking); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", cooking)
}