package repos

import (
	"database/sql"
	"time"

	"github.com/eduardotvn/projeto-api/src/models"
	"github.com/eduardotvn/projeto-api/src/security"
)

type Users struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Users {
	return &Users{db}
}

func (repos Users) Create(user models.User) (sql.Result, error) {
	statement, err := repos.db.Prepare("INSERT INTO Users (created_at, updated_at, deleted_at, name, email, password, admin) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	createdAt := time.Now()
	updatedAt := createdAt
	deletedAt := time.Time{}
	user.Password, err = security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	createdUser, err := statement.Exec(createdAt, updatedAt, deletedAt, user.Name, user.Email, user.Password, user.Admin)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (repos Users) GetAll() ([]models.User, error) {
	results, err := repos.db.Query("SELECT id, name, email, created_at, updated_at FROM Users")
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
			&user.Email,
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
	result, err := repos.db.Query("SELECT id, name, email, created_at, updated_at FROM Users WHERE id = $1", id)
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
			&tempUser.Email,
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

func (repos Users) UpdateUserPasswordById(newPassword, id string) (sql.Result, error) {
	statement, err := repos.db.Prepare("UPDATE Users SET password = $1, updated_at = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	newPassword, err = security.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}

	updatedUserPassword, err := statement.Exec(newPassword, time.Now(), id)
	if err != nil {
		return nil, err
	}

	return updatedUserPassword, nil
}
