package postgres

import (
	"fmt"
	"log"
)

func CreateTables() {
	db, err := DBConnect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, deleted_at TIMESTAMP NOT NULL, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password BYTEA NOT NULL, admin BOOL DEFAULT FALSE)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL PRIMARY KEY, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, title TEXT NOT NULL, content TEXT NOT NULL, user_id INTEGER NOT NULL REFERENCES users(id));")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Necessary Tables checked/created")
}
