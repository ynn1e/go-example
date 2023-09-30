package main

import (
	"fmt"
	"reflect"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user1 := user{
		ID:   1,
		Name: "Tanaka",
		Age:  25,
	}

	rv := reflect.ValueOf(&user1)
	fmt.Printf("value %+v\n", rv)
	fmt.Printf("value kind is %+v\n", rv.Kind())
	elem := rv.Elem()
	fmt.Printf("elem is %+v\n", elem)
	fmt.Printf("elem kind is %+v\n", elem.Kind())
	t := elem.Type()
	fmt.Printf("elem type is %+v\n", t)
	for i := 0; i < elem.NumField(); i++ {
		f := elem.Field(i)
		fmt.Printf("  elem field[%d] is %+v\n", i, f)
		fmt.Printf("  elem field[%d] kind is %+v\n", i, f.Kind())
		fmt.Printf("  elem field[%d] type tag is %+v\n", i, t.Field(i).Tag.Get("json"))
	}
}
