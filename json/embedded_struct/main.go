package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Teacher struct {
	Person
	Subject string `json:"subject"`
}

type Student struct {
	Person
	Score int `json:"score"`
}

func main() {
	teacher := Teacher{
		Person: Person{
			Name: "tanaka",
			Age:  30,
		},
		Subject: "math",
	}
	student := Student{
		Person: Person{
			Name: "suzuki",
			Age:  17,
		},
		Score: 85,
	}
	/*
		teacher: {Person:{Name:tanaka Age:30} Subject:math}
		student: {Person:{Name:suzuki Age:17} Score:85}
	*/
	fmt.Printf("teacher: %+v\n", teacher)
	fmt.Printf("student: %+v\n", student)

	tbyt, err := json.Marshal(teacher)
	if err != nil {
		log.Fatalf("json marshal error: %s", err)
	}
	sbyt, err := json.Marshal(student)
	if err != nil {
		log.Fatalf("json marshal error: %s", err)
	}
	/*
		teacher: {"name":"tanaka","age":30,"subject":"math"}
		teacher: {"name":"suzuki","age":17,"score":85}
	*/
	fmt.Printf("teacher: %s\n", tbyt)
	fmt.Printf("teacher: %s\n", sbyt)
}
