package main

import (
	"fmt"
	"log"
	"time"

	as "github.com/aerospike/aerospike-client-go/v6"
)

func main() {
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	key, err := as.NewKey("test", "myset", "mykey")
	if err != nil {
		log.Fatalln(err)
	}

	bins := as.BinMap{
		"name":  "Tanaka",
		"age":   30,
		"email": "tanaka@example.com",
	}

	wpolisy := as.NewWritePolicy(0, 0)
	if err := client.Put(wpolisy, key, bins); err != nil {
		log.Fatalln(err)
	}

	rpolisy := as.NewPolicy()
	record, err := client.Get(rpolisy, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("--- record ---")
	fmt.Printf("node host: %s\n", record.Node.GetHost())
	fmt.Printf("key namespace: %s\n", record.Key.Namespace())
	fmt.Printf("key value: %s\n", record.Key.Value())
	fmt.Printf("name: %s\n", record.Bins["name"])
	fmt.Printf("age: %d\n", record.Bins["age"])
	fmt.Printf("email: %s\n", record.Bins["email"])
	fmt.Printf("dummy: %v\n", record.Bins["dummy"])
	fmt.Printf("expiration: %ds\n", record.Expiration)

	wpolisy.Expiration = uint32(time.Second * 0)
	ok, err := client.Delete(wpolisy, key)
	if err != nil {
		log.Fatalln(err)
	}

	if !ok {
		log.Fatalf("failed to delete key: %s", key)
	}
}
