package repos

import (
	"fmt"
	"net/mail"
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func (repos Users) ValidateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, ValidationError{Message: "invalid email address"}
	}

	results, err := repos.db.Query("SELECT email FROM Users WHERE email = $1", email)
	if err != nil {

		return false, err
	}
	defer results.Close()

	var Emails struct {
		Email string
	}

	for results.Next() {
		if err = results.Scan(
			&Emails.Email,
		); err != nil {
			return false, err
		}

		if Emails.Email == email {
			return false, ValidationError{Message: "email already exists"}
		}
	}

	return true, nil
}
