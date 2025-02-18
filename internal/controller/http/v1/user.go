package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/auth"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface) {
	r := &userRoutes{t, l}

	h := handler.Group("/users").Use(auth.Auth())
	{
		h.GET("", r.getUsers)
	}
}

type getUsersResponse struct {
	Users []entity.User `json:"users"`
	Total int64         `json:"total"`
}

// @Summary     Get users
// @Description Get users
// @ID          get_users
// @Tags  	    Users
// @Accept      json
// @Produce     json
// @Success     200 {object} getUsersResponse
// @Failure     500 {object} response
// @Router      /v1/users/ [get]
func (r userRoutes) getUsers(context *gin.Context) {
	users, total, err := r.t.Get(context)

	if err != nil {
		context.JSON(500, response{err.Error()})
	}

	context.JSON(200, getUsersResponse{users, total})
}
