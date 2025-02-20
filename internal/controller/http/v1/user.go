package v1

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/auth"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/permission"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
	"net/http"
	"strconv"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface, casbinEnforcer *casbin.Enforcer) {
	r := &userRoutes{t, l}

	h := handler.Group("/users").Use(auth.Auth())
	{
		h.GET("", r.getUsers)
		h.GET("/:id", r.getUser)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "user", "write"), r.update)
	}
}

type getUsersResponse struct {
	Users []entity.User `json:"users"`
	Total int64         `json:"total"`
}

// Get a list of users with pagination.
//
// @Summary      Get Users
// @Description  Fetch a paginated list of users.
// @ID           get_users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200 {object} getUsersResponse "Successful response with user list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/users [get]
func (r userRoutes) getUsers(context *gin.Context) {
	users, total, err := r.t.Get(context)

	if err != nil {
		context.JSON(500, response{err.Error()})
	}

	context.JSON(200, getUsersResponse{users, total})
}

// Get a specific user by ID.
//
// @Summary      Get User
// @Description  Fetch details of a user by their ID.
// @ID           get_user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entity.User "Successful response with user details"
// @Failure      500  {object}  response "Internal server error"
// @Router       /v1/users/{id} [get]
func (r userRoutes) getUser(context *gin.Context) {
	user, err := r.t.GetUser(context)

	if err != nil {
		context.JSON(500, response{err.Error()})
	}

	context.JSON(200, user)
}

// Update user details by ID.
//
// @Summary      Update User
// @Description  Update details of an existing user by providing user ID and update data.
// @ID           update_user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int             true  "User ID"
// @Param        body body      dto.UpdateUserDTO true  "User update data"
// @Success      200  {object}  response "User successfully updated"
// @Failure      400  {object}  response "Invalid request body"
// @Failure      500  {object}  response "Internal server error"
// @Router       /v1/users/{id} [patch]
func (r userRoutes) update(context *gin.Context) {
	var updateUserDTO = &dto.UpdateUserDTO{}
	if err := context.BindJSON(updateUserDTO); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})

		return
	}

	err = r.t.Update(id, updateUserDTO)

	if err != nil {
		context.JSON(500, response{err.Error()})
	}

	context.JSON(200, response{"OK"})
}
