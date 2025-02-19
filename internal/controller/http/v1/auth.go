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

// @Summary     Login
// @Description Login
// @ID          login
// @Tags  	    Auth
// @Accept      json
// @Produce     json
// @Success     303 {object} loginResponse
// @Failure     500 {object} response
// @Router      /v1/login [POST]
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

// @Summary     Register
// @Description Register
// @ID          register
// @Tags  	    Auth
// @Accept      json
// @Produce     json
// @Success     200 {object} loginResponse
// @Failure     500 {object} response
// @Router      /v1/register [POST]
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
