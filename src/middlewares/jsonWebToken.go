package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = os.Getenv("JWT_KEY")

type UserClaims struct {
	UserID          uint   `json:"user_id"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	UserAdmin       bool   `json:"user_admin"`
	Authorized      bool   `json:"authorized"`
	AdminAuthorized bool   `json:"admin_authorized"`
	jwt.StandardClaims
}

type UserLoginInformation struct {
	ID    uint
	Name  string
	Email string
	Admin bool
}

func GenerateWebToken(user UserLoginInformation) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &UserClaims{
		UserID:    user.ID,
		UserName:  user.Name,
		UserEmail: user.Email,
		UserAdmin: user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	claims.Authorized = true

	if user.Admin {
		claims.AdminAuthorized = true
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(signingKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.New("could not sign token")
	}

	return tokenString, nil
}
