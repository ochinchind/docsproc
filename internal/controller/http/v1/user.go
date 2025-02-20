package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/auth"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/permission"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
	"log"
	"net/http"
	"strconv"
	"time"
)

type userRoutes struct {
	t  usecase.User
	l  logger.Interface
	rd *redis.Client
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &userRoutes{t, l, rd}

	h := handler.Group("/users").Use(auth.Auth(rd))
	{
		h.GET("", r.getUsers)
		h.GET("/:id", r.getUser)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "user", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "user", "write"), r.delete)
		h.PATCH("/myprofile", r.updateProfile)
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
		return
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
		return
	}

	r.blacklistTokens(context.Param("id"))

	context.JSON(200, response{"OK"})
}

// Update user profile.
//
// @Summary      Update Profile
// @Description  Update profile details of the currently authenticated user.
// @ID           update_profile
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body      dto.UpdateProfileDTO true  "User update data"
// @Success      200  {object}  response "User profile successfully updated"
// @Failure      400  {object}  response "Invalid request body"
// @Failure      500  {object}  response "Internal server error"
// @Router       /v1/users/myprofile [patch]
func (r userRoutes) updateProfile(context *gin.Context) {
	var updateProfileDTO = &dto.UpdateProfileDTO{}
	if err := context.BindJSON(updateProfileDTO); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	id := context.GetInt("auth_user_id")

	updateUserDTO := &dto.UpdateUserDTO{
		Username: updateProfileDTO.Username,
		Name:     updateProfileDTO.Name,
		Surname:  updateProfileDTO.Surname,
		Password: updateProfileDTO.Password,
		Phone:    updateProfileDTO.Phone,
	}

	err := r.t.Update(id, updateUserDTO)

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	context.JSON(200, response{"OK"})
}

// Delete a user by ID.
//
// @Summary      Delete User
// @Description  Delete an existing user by providing user ID.
// @ID           delete_user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int true  "User ID"
// @Success      200  {object}  response "User successfully deleted"
// @Failure      500  {object}  response "Internal server error"
// @Router       /v1/users/{id} [delete]
func (r userRoutes) delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})

		return
	}

	err = r.t.Delete(id)

	if err != nil {
		context.JSON(500, response{err.Error()})

		return
	}

	r.blacklistTokens(context.Param("id"))

	context.JSON(200, response{"OK"})
}

func (r userRoutes) blacklistTokens(id string) {
	go func() {
		cacheKey := fmt.Sprintf("user_id_%s", id)

		// Fetch existing tokens from Redis
		result, err := r.rd.Get(cacheKey).Result()
		var tokens []string

		if !errors.Is(err, redis.Nil) {
			// Deserialize JSON array from Redis
			err = json.Unmarshal([]byte(result), &tokens)
			if err != nil {
				log.Println("Error unmarshalling JSON:", err)
				return
			}

			// Check if the token already exists
			for _, t := range tokens {
				blacklistKey := fmt.Sprintf("blacklist_%s", t)
				err := r.rd.Set(blacklistKey, "true", 24*time.Hour).Err()
				if err != nil {
					log.Println("Error blacklisting token:", err)
					return
				}
			}
		}
	}()
}
