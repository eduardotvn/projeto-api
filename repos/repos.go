package repos

import (
	"database/sql"
	"time"

	"github.com/eduardotvn/projeto-api/src/models"
)

type Users struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Users {
	return &Users{db}
}

func (repos Users) Create(user models.User) (sql.Result, error) {
	statement, err := repos.db.Prepare("INSERT INTO Users (id, created_at, updated_at, deleted_at, name, password, admin) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	createdAt := time.Now()
	updatedAt := createdAt
	deletedAt := time.Time{}

	createdUser, err := statement.Exec(user.ID, createdAt, updatedAt, deletedAt, user.Name, user.Password, user.Admin)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (repos Users) GetAll() ([]models.User, error) {
	results, err := repos.db.Query("SELECT id, name, created_at, updated_at FROM Users")
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var allUsers []models.User

	for results.Next() {
		var user models.User

		if err = results.Scan(
			&user.ID,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		allUsers = append(allUsers, user)
	}

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}

func (repos Users) GetUserByID(id string) (models.User, error) {
	result, err := repos.db.Query("SELECT id, name, created_at, updated_at FROM Users WHERE id = $1", id)
	if err != nil {
		return models.User{}, err
	}
	defer result.Close()

	var user models.User

	for result.Next() {

		var tempUser models.User

		if err = result.Scan(
			&tempUser.ID,
			&tempUser.Name,
			&tempUser.CreatedAt,
			&tempUser.UpdatedAt,
		); err != nil {
			return models.User{}, err
		}

		user = tempUser
	}

	return user, nil
}

func (repos Users) DeleteUserById(id string) error {
	statement, err := repos.db.Prepare("DELETE FROM Users WHERE ID = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
