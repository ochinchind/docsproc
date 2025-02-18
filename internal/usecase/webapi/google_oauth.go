package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/config"
	"github.com/ochinchind/docsproc/internal/entity"
	"io"
	"net/http"
)

// GoogleOAuthWebApi -.
type GoogleOAuthWebApi struct {
}

// New -.
func New() *GoogleOAuthWebApi {
	return &GoogleOAuthWebApi{}
}

// FetchUserFromGoogle -.
func (t *GoogleOAuthWebApi) FetchUserFromGoogle(context *gin.Context) (*entity.GoogleCallbackResponse, error) {
	state := context.Query("state")
	if state != "kno31jn12j4nk324nj" {
		return nil, fmt.Errorf("GoogleOAuthWebApi - FetchUserFromGoogle - context.Query")
	}

	code := context.Query("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Request.Context(), code)
	if err != nil {
		return nil, fmt.Errorf("GoogleOAuthWebApi - FetchUserFromGoogle - googlecon.Exchange: %w", err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("GoogleOAuthWebApi - FetchUserFromGoogle - http.Get: %w", err)
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GoogleOAuthWebApi - FetchUserFromGoogle - io.ReadAll: %w", err)
	}

	var userInfo entity.GoogleCallbackResponse
	if err := json.Unmarshal(userData, &userInfo); err != nil {
		return nil, fmt.Errorf("GoogleOAuthWebApi - FetchUserFromGoogle - json.Unmarshal: %w", err)
	}

	return &userInfo, nil
}
