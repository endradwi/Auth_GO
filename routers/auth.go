package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func AuthsRouter(router *gin.RouterGroup) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUsers)
}
