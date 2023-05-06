package repos

import (
	"database/sql"
	"time"

	"github.com/eduardotvn/projeto-api/src/models"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repos Posts) InsertPost(post models.Post, id uint) (sql.Result, error) {
	statement, err := repos.db.Prepare("INSERT INTO Posts (created_at, updated_at, title, content, user_id) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	createdAt := time.Now()
	updatedAt := createdAt

	createPost, err := statement.Exec(createdAt, updatedAt, post.Title, post.Content, id)
	if err != nil {
		return nil, err
	}

	return createPost, nil
}
