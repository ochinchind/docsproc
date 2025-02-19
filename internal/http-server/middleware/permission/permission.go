package permission

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Permission(casbinEnforcer *casbin.Enforcer, permission string) gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("auth_user_role")

		if !exists {
			context.JSON(500, gin.H{"error": "user role not found"})
			context.Abort()
			return
		}

		path := context.Request.URL.Path
		fmt.Println(path, role, permission)

		exists, err := casbinEnforcer.Enforce(role, path, permission)

		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		if !exists {
			context.JSON(403, gin.H{"error": "forbidden"})
			context.Abort()
			return
		}

		context.Next()
	}
}
