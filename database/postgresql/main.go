package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	db *sql.DB
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var (
		cfg      Config
		user     = os.Getenv("USER")
		dbname   = os.Getenv("DBNAME")
		password = os.Getenv("PASSWORD")
	)

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cfg.db = db

	http.HandleFunc("/users", cfg.handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen and serve error: %s", err)
	}
}

func (cfg *Config) handler(w http.ResponseWriter, r *http.Request) {
	rows, err := cfg.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("query error: %s", err)

		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var (
			id, age int
			name    string
		)

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("scan error: %s", err)

			return
		}

		users = append(users, User{
			ID:   id,
			Name: name,
			Age:  age,
		})
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json marshal error: %s", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(data)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("write error: %s", err)

		return
	}
}
