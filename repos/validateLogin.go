package repos

import (
	"fmt"

	"github.com/eduardotvn/projeto-api/src/middlewares"
	"github.com/eduardotvn/projeto-api/src/security"
)

func (repos Users) ValidateLogin(email, password string) (middlewares.UserLoginInformation, error) {
	var user middlewares.UserLoginInformation
	result, err := repos.db.Query("SELECT id, name, email, password, admin FROM Users WHERE email = $1", email)
	if err != nil {
		return middlewares.UserLoginInformation{}, err
	}
	defer result.Close()

	var hashedPassword string

	for result.Next() {

		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&hashedPassword,
			&user.Admin,
		); err != nil {
			return middlewares.UserLoginInformation{}, err
		}
	}

	_, err = security.ValidatePassword(password, hashedPassword)
	if err != nil {
		return middlewares.UserLoginInformation{}, fmt.Errorf("Wrong Password")
	} else {
		return user, nil
	}
}
