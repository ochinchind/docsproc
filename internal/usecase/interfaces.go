package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
)

type (
	// GoogleOAuth -.
	GoogleOAuth interface {
		GoogleLogin() string
		GoogleCallback(context *gin.Context) (string, error)
	}

	// UserRepo -.
	UserRepo interface {
		GetByEmail(email string) (entity.User, error)
		Create(context *gin.Context, user entity.User) error
		Get(context *gin.Context) ([]entity.User, int64, error)
	}

	// GoogleOAuthWebApi -.
	GoogleOAuthWebApi interface {
		FetchUserFromGoogle(context *gin.Context) (*entity.GoogleCallbackResponse, error)
	}

	// JWT -.
	JWT interface {
		GenerateJWT(email, username string) (string, error)
		ValidateToken(tokenString string) error
	}

	User interface {
		Get(context *gin.Context) ([]entity.User, int64, error)
	}
)
