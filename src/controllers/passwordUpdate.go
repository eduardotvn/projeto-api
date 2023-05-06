package controllers

import (
	"github.com/eduardotvn/projeto-api/repos"
	"github.com/gin-gonic/gin"
)

func LoadPasswordChangePage(c *gin.Context) {
	id := c.Param("id")

	id = repos.GenerateToken(id)

	//Load password changer
}
