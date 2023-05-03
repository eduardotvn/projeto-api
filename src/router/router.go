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
		main.POST("/register", controllers.CreateUser)
		main.GET("/getUser/:id", controllers.GetUser)
		main.PUT("/updateUser/:id", controllers.UpdateUser)
		admin := main.Group("admin")
		{
			admin.GET("getAll", controllers.GetAllUsers)
			admin.DELETE("/deleteUser/:id", controllers.DeleteUser)

		}
	}

	return r
}
