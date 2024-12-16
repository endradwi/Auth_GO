package routers

import (
	"test/controllers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidationToken())
	router.GET("", controllers.GetAllUsers)
	router.GET("/:id", controllers.GetUserById)
	router.POST("", controllers.GetUserById)
	router.PATCH("/:id", controllers.EditUser)
	router.DELETE("/:id", controllers.DeleteUser)
}
