package router

import (
	"net/http"

	"github.com/eduardotvn/projeto-api/src/controllers"
	"github.com/eduardotvn/projeto-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

// EM CONSTRUÇÃO E EVOLUÇÃO
func CreateRouter(r *gin.Engine) *gin.Engine {
	main := r.Group("")
	{
		main.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
		main.POST("/register", controllers.CreateUser)
		main.POST("/login", controllers.LoginUser)
		user := main.Group("user")
		user.Use(middlewares.AuthMiddleware())
		{
			user.GET("/getUser/:id", controllers.GetUser)
			user.PUT("/updateUser/:id", controllers.UpdateUserPassword)
			user.POST("/createPost", controllers.CreatePost)
		}
		admin := main.Group("admin")
		admin.Use(middlewares.AuthMiddleware())
		{
			admin.GET("getAll", controllers.GetAllUsers)
			admin.DELETE("/deleteUser/:id", controllers.DeleteUser)

		}
	}

	return r
}
