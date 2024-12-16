package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	UsersRouter(router.Group("/users"))
	MovieRouter(router.Group("/movies"))
	AuthsRouter(router.Group("/auth"))
}
