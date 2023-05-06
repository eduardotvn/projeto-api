package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eduardotvn/projeto-api/repos"
	"github.com/eduardotvn/projeto-api/response"
	"github.com/eduardotvn/projeto-api/src/models"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	bodyRequest, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		response.Error(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repos := repos.NewRepo(db)
	_, err = repos.ValidateEmail(user.Email)
	if err != nil {
		fmt.Println(err)
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}
	_, err = repos.Create(user)
	if err != nil {
		fmt.Println(err)
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(c.Writer, http.StatusCreated, user)
}

/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/

func GetAllUsers(c *gin.Context) {

	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repos := repos.NewRepo(db)
	users, err := repos.GetAll()
	if err != nil {
		fmt.Println(err)
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(c.Writer, http.StatusAccepted, users)
}

/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/

func GetUser(c *gin.Context) {
	id := c.Param("id")

	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repos := repos.NewRepo(db)

	user, err := repos.GetUserByID(id)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(c.Writer, http.StatusOK, user)
}

/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/

func UpdateUserPassword(c *gin.Context) {

	bodyRequest, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		response.Error(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	id := c.Param("id")

	var newPassword struct {
		Password string
	}

	if err = json.Unmarshal(bodyRequest, &newPassword); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repos := repos.NewRepo(db)
	updatedUser, err := repos.UpdateUserPasswordById(newPassword.Password, id)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(c.Writer, http.StatusOK, updatedUser)
}

/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	db, err := postgres.DBConnect()
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repos := repos.NewRepo(db)
	if err := repos.DeleteUserById(id); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(c.Writer, 200, nil)
}

/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
/*----------------------------------------------------------------------------------------------------*/
