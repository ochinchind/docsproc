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

// Google login initiates OAuth authentication with Google.
//
// @Summary      Google Login
// @Description  Redirects the user to Google's OAuth authentication page.
// @ID           google_login
// @Tags         GoogleOAuth
// @Accept       json
// @Produce      json
// @Success      303 "Redirect to Google login"
// @Failure      500 {object} response "Internal server error"
// @Router       /google_login [get]
func (r googleOAuthRoutes) googleLogin(context *gin.Context) {
	url := r.t.GoogleLogin()

	fmt.Println(url)

	context.Redirect(http.StatusSeeOther, url)
}

type googleCallbackResponse struct {
	Token string `json:"token"`
}

// Google callback handles the OAuth callback from Google.
//
// @Summary      Google OAuth Callback
// @Description  Handles Google's OAuth callback, processes authentication, and returns an access token.
// @ID           google_callback
// @Tags         GoogleOAuth
// @Accept       json
// @Produce      json
// @Success      200 {object} googleCallbackResponse "Successfully authenticated"
// @Failure      500 {object} response "Internal server error"
// @Router       /google_callback [get]
func (r googleOAuthRoutes) googleCallback(context *gin.Context) {
	token, err := r.t.GoogleCallback(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
	}

	context.JSON(http.StatusOK, googleCallbackResponse{token})
}
