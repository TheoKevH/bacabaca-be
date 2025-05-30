// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID        pgtype.UUID
	Title     string
	Slug      string
	Content   string
	AuthorID  pgtype.UUID
	CreatedAt pgtype.Timestamptz
}

type User struct {
	ID        pgtype.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamptz
}
