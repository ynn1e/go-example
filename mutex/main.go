package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu   = sync.Mutex{}
	rwmu = sync.RWMutex{}
	data = 0
	wg   = sync.WaitGroup{}
)

/*
$ go test -bench .
mu read: 0
mu read: 0
mu read: 0
mu write: 1
goos: darwin
goarch: arm64
BenchmarkUseMu-8     	       1	6004640084 ns/op
rwmu read: 1
rwmu read: 1
rwmu read: 1
rwmu write: 2
BenchmarkUseRWMu-8   	       1	4002299708 ns/op
PASS
ok  	_/Users/shunya.inoue/go-example/mutex	10.909s
*/

func main() {
	UseMu()

	UseRWMu()
}

func UseMu() {
	wg.Add(4)

	go muRead()
	go muRead()
	go muWrite()
	go muRead()

	wg.Wait()
}

func UseRWMu() {
	wg.Add(4)

	go rwmuRead()
	go rwmuRead()
	go rwmuWrite()
	go rwmuRead()

	wg.Wait()
}

func muRead() {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	time.Sleep(time.Second)
	fmt.Printf("mu read: %d\n", data)
}

func muWrite() {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	time.Sleep(3 * time.Second)
	data++
	fmt.Printf("mu write: %d\n", data)
}

func rwmuRead() {
	defer wg.Done()

	rwmu.RLock()
	defer rwmu.RUnlock()

	time.Sleep(time.Second)
	fmt.Printf("rwmu read: %d\n", data)
}

func rwmuWrite() {
	defer wg.Done()

	rwmu.Lock()
	defer rwmu.Unlock()

	time.Sleep(3 * time.Second)
	data++
	fmt.Printf("rwmu write: %d\n", data)
}
