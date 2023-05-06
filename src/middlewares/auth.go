package middlewares

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var signingKeyBytes []byte

func init() {
	var err error
	signingKeyBytes, err = base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		panic(fmt.Sprintf("failed to decode JWT_KEY: %v", err))
	}
}

func parseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return signingKeyBytes, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := parseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !claims.Authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not authorized as user",
			})
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := parseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !claims.Authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not authorized as user",
			})
			return
		}

		if !claims.AdminAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "not an admin",
			})
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
