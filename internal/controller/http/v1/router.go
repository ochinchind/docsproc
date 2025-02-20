// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis"
	"github.com/ochinchind/docsproc/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/ochinchind/docsproc/docs"
	"github.com/ochinchind/docsproc/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, s *usecase.Services, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	g := handler.Group("")
	{
		newGoogleOAuthRoutesRoutes(g, s.GoogleOAuth, l, casbinEnforcer)
		newAuthRoutes(g, s.Auth, l)
	}

	h := handler.Group("/v1")
	{
		newUserRoutes(h, s.User, l, casbinEnforcer, rd)
		newSpecialtyRoutes(h, s.Specialty, l, casbinEnforcer, rd)
		newQualificationRoutes(h, s.Qualification, l, casbinEnforcer, rd)
		newDisciplineRoutes(h, s.Discipline, l, casbinEnforcer, rd)
	}
}
