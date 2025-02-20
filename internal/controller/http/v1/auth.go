package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
	"net/http"
)

type authRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
	r := &authRoutes{t, l}

	handler.POST("/login", r.login)
	handler.POST("/register", r.register)
}

type loginResponse struct {
	Token string `json:"token"`
}

// Login handles user authentication and token generation.
//
// @Summary      User Login
// @Description  Authenticates a user and returns an access token.
// @ID           login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginDTO true "Login request body"
// @Success      200 {object} loginResponse "Successfully authenticated"
// @Failure      400 {object} response "Invalid credentials"
// @Failure      406 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /login [post]
func (r authRoutes) login(context *gin.Context) {
	var loginDTO = &dto.LoginDTO{}
	if err := context.BindJSON(loginDTO); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	token, err := r.t.Login(loginDTO)

	if err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, loginResponse{token})
}

// Register handles new user registration.
//
// @Summary      User Registration
// @Description  Registers a new user and returns an access token.
// @ID           register
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterDTO true "Register request body"
// @Success      200 {object} loginResponse "Successfully registered"
// @Failure      400 {object} response "Invalid request data"
// @Failure      406 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /register [post]
func (r authRoutes) register(context *gin.Context) {
	var registerDTO = &dto.RegisterDTO{}
	if err := context.BindJSON(registerDTO); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	token, err := r.t.Register(registerDTO)

	if err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, loginResponse{token})
}
