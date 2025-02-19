package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/config"
	"github.com/ochinchind/docsproc/internal/entity"
)

// GoogleOAuthUseCase -.
type GoogleOAuthUseCase struct {
	googleOAuthWebApi GoogleOAuthWebApi
	userRepo          UserRepo
}

// NewGoogleOAuthUseCase -.
func NewGoogleOAuthUseCase(googleOAuthWebApi GoogleOAuthWebApi, userRepo UserRepo) *GoogleOAuthUseCase {
	return &GoogleOAuthUseCase{
		googleOAuthWebApi: googleOAuthWebApi,
		userRepo:          userRepo,
	}
}

// GoogleLogin Google login
func (uc *GoogleOAuthUseCase) GoogleLogin() string {
	loginConfig := config.Cfg.GoogleLoginConfig
	url := loginConfig.AuthCodeURL("kno31jn12j4nk324nj")

	return url
}

// GoogleCallback Google callback
func (uc *GoogleOAuthUseCase) GoogleCallback(context *gin.Context) (string, error) {
	userInfo, err := uc.googleOAuthWebApi.FetchUserFromGoogle(context)

	if err != nil {
		return "", fmt.Errorf("GoogleOAuthUseCase - GoogleCallback - uc.GoogleOAuthWebApi.FetchUserFromGoogle: %w", err)
	}

	user, err := uc.userRepo.GetByEmail(userInfo.Email)

	if err != nil {
		return "", fmt.Errorf("GoogleOAuthUseCase - GoogleCallback - uc.UserRepo.GetByEmail: %w", err)
	}

	if user != (entity.User{}) {
		tokenString, err := GenerateJWT(user.Email, user.Name, user.Role)
		if err != nil {
			return "", fmt.Errorf("GoogleOAuthUseCase - GoogleCallback - GenerateJWT: %w", err)
		}

		return tokenString, nil
	}

	user = entity.User{
		Email:    userInfo.Email,
		Username: userInfo.Email,
		Name:     userInfo.GivenName,
		Role:     "user",
		Surname:  userInfo.FamilyName,
		Picture:  userInfo.Picture,
	}

	err = uc.userRepo.Create(context, user)

	if err != nil {
		return "", fmt.Errorf("GoogleOAuthUseCase - GoogleCallback - uc.UserRepo.Create: %w", err)
	}

	tokenString, err := GenerateJWT(user.Email, user.Name, user.Role)
	if err != nil {
		return "", fmt.Errorf("GoogleOAuthUseCase - GoogleCallback - GenerateJWT: %w", err)
	}

	return tokenString, nil
}
