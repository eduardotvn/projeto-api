package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eduardotvn/projeto-api/repos"
	"github.com/eduardotvn/projeto-api/response"
	"github.com/eduardotvn/projeto-api/src/postgres"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

	//APENAS PARA TESTE, HASH SERÁ ADICIONADO AO PASSWORD PARA PROTEÇÃO DO USUÁRIO

	bodyRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c.Writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user struct {
		Name     string
		Password string
	}
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
	loginResult, err := repos.ValidateLogin(user.Name, user.Password)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, err)
		return
	}

	response.JSON(c.Writer, http.StatusOK, loginResult)
}
