// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package users

import (
	"database/sql"
)

type User struct {
	ID   int32         `json:"id"`
	Name string        `json:"name"`
	Age  sql.NullInt32 `json:"age"`
}