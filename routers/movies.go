package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	// router.Use(middlewares.ValidationToken())
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMoviesById)
	router.POST("", controllers.SaveMovies)
	// router.PATCH("/:id", controllers.EditMovie)
	router.DELETE("/:id", controllers.DeleteMovie)
}
