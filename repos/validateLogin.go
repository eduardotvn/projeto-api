package repos

import "github.com/eduardotvn/projeto-api/src/security"

func (repos Users) ValidateLogin(email, password string) (bool, error) {

	result, err := repos.db.Query("SELECT email, password FROM Users WHERE name = $1 AND password = $2", email, password)
	if err != nil {
		return false, err
	}
	defer result.Close()

	var temp struct {
		Email    string
		Password string
	}

	for result.Next() {

		if err = result.Scan(
			&temp.Email,
			&temp.Password,
		); err != nil {
			return false, err
		}
	}

	password, err = security.HashPassword(password)
	if err != nil {
		return false, nil
	}

	if temp.Email == email && temp.Password == password {
		return true, nil
	} else {
		return false, nil
	}
}
