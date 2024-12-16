package routers

import (
	"test/controllers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	// router.Use(middlewares.ValidationToken())
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMoviesById)
	router.POST("", middlewares.ValidationToken(), controllers.SaveMovies)
	router.PATCH("/:id", middlewares.ValidationToken(), controllers.EditMovie)
	router.DELETE("/:id", middlewares.ValidationToken(), controllers.DeleteMovie)
}
