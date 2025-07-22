package middleware

import (
	"net/http"
	"task-manager/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")

		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No Authentication header provided"})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)
		ctx.Set("first_name", claims.FirstName)
		ctx.Set("last_name", claims.LastName)
		ctx.Set("user_id", claims.Uid)
		ctx.Set("user_type", claims.UserType)

		ctx.Next()
	}
}
