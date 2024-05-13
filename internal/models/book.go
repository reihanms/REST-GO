package models

import "time"

type Book struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rating      int
	Created_at  time.Time
	Updated_at  time.Time
}
