package repos

func (repos Users) ValidateLogin(name, password string) (bool, error) {
	result, err := repos.db.Query("SELECT name, password FROM Users WHERE name = $1 AND password = $2", name, password)
	if err != nil {
		return false, err
	}
	defer result.Close()

	var temp struct {
		Name     string
		Password string
	}

	for result.Next() {

		if err = result.Scan(
			&temp.Name,
			&temp.Password,
		); err != nil {
			return false, err
		}
	}

	if temp.Name == name && temp.Password == password {
		return true, nil
	} else {
		return false, nil
	}
}
