package main

import (
	"os"
	"sync"
	"time"
)

func main() {
	switch os.Getenv("RACE") {
	case "1":
		raceCondition1()
	case "2":
		raceCondition2()
	default:
		noRaceCondition()
	}
}

/*
$ RACE=1 go run -race main.go
==================
WARNING: DATA RACE
...
==================
World
World
Found 1 data race(s)
exit status 66
*/
func raceCondition1() {
	msg := "Hello"
	go func() {
		println(msg)
	}()
	msg = "World"

	time.Sleep(time.Second)
	println(msg)
}

/*
$ RACE=2 go run -race main.go
Hello
==================
WARNING: DATA RACE
...
==================
World
Found 1 data race(s)
exit status 66
*/
func raceCondition2() {
	msg := "Hello"
	go func() {
		println(msg)
	}()
	time.Sleep(time.Second)
	msg = "World"

	time.Sleep(time.Second)
	println(msg)
}

/*
$ go run -race main.go
Hello
World
*/
func noRaceCondition() {
	msg := "Hello"
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		println(msg)
	}()
	wg.Wait()
	msg = "World"

	time.Sleep(time.Second)
	println(msg)
}
