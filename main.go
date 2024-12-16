package main

import (
	"test/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// router.Use(middlewares.SetHTMLHeader())
	routers.Routers(router)

	router.Run("localhost:8888")
}
