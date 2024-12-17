package middlewares

import (
	"net/http"
	"strings"
	"test/controllers"
	"test/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func ValidationToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		head := ctx.GetHeader("Authorization")

		if head == "" {
			ctx.JSON(http.StatusNotFound, controllers.Response{
				Success: false,
				Message: "Token not found",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]

		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})

		out := jwt.Claims{}

		err := tok.Claims(lib.JWT_SECRET, &out)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
		}
		// if head != "true" {
		// 	ctx.JSON(http.StatusUnauthorized, controllers.Response{
		// 		Success: false,
		// 		Message: "Unauthorized",
		// 	})
		// 	ctx.Abort()
		// }
		ctx.Next()
	}
}
