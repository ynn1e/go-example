package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type env struct {
	debug bool
}

func (e *env) handler(w http.ResponseWriter, _ *http.Request) {
	if e.debug {
		if _, err := io.WriteString(w, "debug mode\n"); err != nil {
			log.Fatalf("handler: %s", err)
		}
	}

	if _, err := io.WriteString(w, "👋\n"); err != nil {
		log.Fatalf("handler: %s", err)
	}
}

func healthcheck(w http.ResponseWriter, _ *http.Request) {
	if _, err := io.WriteString(w, "ok\n"); err != nil {
		log.Fatalf("healthcheck: %s", err)
	}
}

/*
$ curl "localhost:8080/healthcheck"
ok
$ curl "localhost:8080"
404 page not found
$ curl "localhost:8080/aaa"
404 page not found
$ curl "localhost:8080/hello"
debug mode
👋
*/

func main() {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debug = true
	}

	env := env{debug: debug}

	http.HandleFunc("/", http.NotFound)
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/hello", env.handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
