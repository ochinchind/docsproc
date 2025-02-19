package v1

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
	"net/http"
)

type googleOAuthRoutes struct {
	t usecase.GoogleOAuth
	l logger.Interface
}

func newGoogleOAuthRoutesRoutes(handler *gin.RouterGroup, t usecase.GoogleOAuth, l logger.Interface, casbinEnforcer *casbin.Enforcer) {
	r := &googleOAuthRoutes{t, l}

	handler.GET("/google_login", r.googleLogin)
	handler.GET("/google_callback", r.googleCallback)
}

// @Summary     Google login
// @Description Google login
// @ID          google_login
// @Tags  	    GoogleOAuth
// @Accept      json
// @Produce     json
// @Success     303 {object} googleLoginResponse
// @Failure     500 {object} response
// @Router      /google_login [get]
func (r googleOAuthRoutes) googleLogin(context *gin.Context) {
	url := r.t.GoogleLogin()

	fmt.Println(url)

	context.Redirect(http.StatusSeeOther, url)
}

type googleCallbackResponse struct {
	Token string `json:"token"`
}

// @Summary     Google callback
// @Description Google callback
// @ID          google_callback
// @Tags  	    GoogleOAuth
// @Accept      json
// @Produce     json
// @Success     200 {object} googleCallbackResponse
// @Failure     500 {object} response
// @Router      /google_callback [get]
func (r googleOAuthRoutes) googleCallback(context *gin.Context) {
	token, err := r.t.GoogleCallback(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
	}

	context.JSON(http.StatusOK, googleCallbackResponse{token})
}
