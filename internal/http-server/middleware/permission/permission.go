package permission

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Permission(casbinEnforcer *casbin.Enforcer, obj string, permission string) gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("auth_user_role")

		if !exists {
			context.JSON(500, gin.H{"error": "user role not found"})
			context.Abort()
			return
		}

		exists, err := casbinEnforcer.Enforce(role, obj, permission)

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
