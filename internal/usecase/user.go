package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	userRepo UserRepo
}

// NewUserUseCase -.
func NewUserUseCase(userRepo UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// Get -.
func (uc *UserUseCase) Get(context *gin.Context) ([]entity.User, int64, error) {
	users, total, err := uc.userRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
