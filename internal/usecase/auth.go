package usecase

import (
	"fmt"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/utils"
)

// AuthUseCase -.
type AuthUseCase struct {
	userRepo UserRepo
}

// NewAuthUseCase -.
func NewAuthUseCase(userRepo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

// Login login
func (uc *AuthUseCase) Login(dto *dto.LoginDTO) (string, error) {
	user, err := uc.userRepo.GetByUsername(dto.Username)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - uc.UserRepo.GetByUsername: %w", err)
	}

	if user == (entity.User{}) {
		return "", fmt.Errorf("user not found")
	}

	if utils.VerifyPassword(dto.Password, user.Password) {
		return "", fmt.Errorf("wrong password")
	}

	tokenString, err := GenerateJWT(user.Email, user.Name, user.Role)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - GenerateJWT - uc.JWT.GenerateJWT: %w", err)
	}

	return tokenString, nil
}

// Register register
func (uc *AuthUseCase) Register(dto *dto.RegisterDTO) (string, error) {
	user, err := uc.userRepo.GetByUsernameOrEmail(dto.Username, dto.Email)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - uc.UserRepo.GetByUsername: %w", err)
	}

	if user.Email == dto.Email {
		return "", fmt.Errorf("email already exists")
	}

	if user.Username == dto.Username {
		return "", fmt.Errorf("username already exists")
	}

	hash, err := utils.HashPassword(dto.Password)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - HashPassword: %w", err)
	}

	user = entity.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hash,
		Name:     dto.Name,
		Surname:  dto.Surname,
		Phone:    dto.Phone,
		Role:     dto.Role,
	}

	err = uc.userRepo.Create(user)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - uc.UserRepo.Create: %w", err)
	}

	tokenString, err := GenerateJWT(user.Email, user.Name, user.Role)

	if err != nil {
		return "", fmt.Errorf("AuthUseCase - GenerateJWT - uc.JWT.GenerateJWT: %w", err)
	}

	return tokenString, nil
}
