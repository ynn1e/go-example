package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"example.com/users"
	_ "github.com/lib/pq"
)

var (
	bNoRowsJSON = []byte("{}")
	qUsers      *users.Queries
)

func main() {
	var (
		user     = os.Getenv("USER")
		dbname   = os.Getenv("DBNAME")
		password = os.Getenv("PASSWORD")
	)

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	qUsers = users.New(db)

	http.HandleFunc("/users", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("listen and serve error: %s", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		data   = bNoRowsJSON
		err    error
		values = r.URL.Query()
		vID    = values.Get("id")
	)

	switch {
	case vID == "":
		if data, err = getJSONUsers(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("get json users error: %s", err)

			return
		}
	default:
		id, err := strconv.Atoi(vID)
		if err != nil {
			break
		}

		if data, err = getJSONUser(r.Context(), int32(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("get json user(id=%d) error: %s", id, err)

			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(data)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("write error: %s", err)

		return
	}
}

func getJSONUsers(ctx context.Context) ([]byte, error) {
	usrs, err := qUsers.ListUsers(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return bNoRowsJSON, nil
	}

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(usrs, "", "  ")
}

func getJSONUser(ctx context.Context, id int32) ([]byte, error) {
	usr, err := qUsers.GetUser(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return bNoRowsJSON, nil
	}

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(usr, "", "  ")
}
