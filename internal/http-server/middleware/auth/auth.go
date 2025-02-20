package auth

import (
	"github.com/ochinchind/docsproc/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "dto does not contain an access token"})
			context.Abort()
			return
		}

		// Check if it starts with "Bearer " and remove it
		const bearerPrefix = "Bearer "
		if len(tokenString) > len(bearerPrefix) && tokenString[:len(bearerPrefix)] == bearerPrefix {
			tokenString = tokenString[len(bearerPrefix):] // Extract the actual token
		}

		claims, err := usecase.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("auth_user_role", claims.Role)
		context.Set("auth_user_email", claims.Email)

		context.Next()
	}
}
