package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
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
		GetByUsernameOrEmail(username, email string) (entity.User, error)
		GetByEmail(email string) (entity.User, error)
		GetByUsername(username string) (entity.User, error)
		Create(user *entity.User) error
		Update(user *entity.User) error
		Get(context *gin.Context) ([]entity.User, int64, error)
		GetByID(int) (entity.User, error)
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
		GetUser(context *gin.Context) (*entity.User, error)
		Update(id int, dto *dto.UpdateUserDTO) error
	}

	Auth interface {
		Login(dto *dto.LoginDTO) (string, error)
		Register(dto *dto.RegisterDTO) (string, error)
	}
)
