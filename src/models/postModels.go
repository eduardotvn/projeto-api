package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" primaryKey:"true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Admin     bool      `json:"admin"`
}
