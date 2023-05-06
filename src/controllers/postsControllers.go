package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eduardotvn/projeto-api/repos"
	"github.com/eduardotvn/projeto-api/response"
	"github.com/eduardotvn/projeto-api/src/middlewares"
	"github.com/eduardotvn/projeto-api/src/models"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	bodyRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "missing authorization header",
		})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	claims, err := middlewares.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	repos := repos.NewPostRepo(db)
	newPost, err := repos.InsertPost(post, claims.UserID)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(c.Writer, http.StatusOK, newPost)
}
