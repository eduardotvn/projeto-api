package router

import (
	"net/http"

	"github.com/eduardotvn/projeto-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter(r *gin.Engine) *gin.Engine {
	main := r.Group("")
	{
		main.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
		main.POST("/createUser", controllers.CreateUser)
		main.GET("/getAll", controllers.GetAllUsers)
		main.GET("/getUser/:id", controllers.GetUser)
		main.PUT("/updateUser/:id", controllers.UpdateUser)
		main.DELETE("/deleteUser/:id", controllers.DeleteUser)
	}

	return r
}
