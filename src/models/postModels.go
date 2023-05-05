package models

import (
	"time"
)

// ADICIONAR IMPLEMENTAÇÃO AUTOMÁTICA DO VALOR DO ID
type User struct {
	ID        uint      `json:"id" primaryKey:"true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Admin     bool      `json:"admin"`
}
