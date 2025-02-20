package usecase

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/utils"
	"gorm.io/gorm"
	"strconv"
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

// GetUser -.
func (uc *UserUseCase) GetUser(context *gin.Context) (*entity.User, error) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (uc *UserUseCase) Update(id int, dto *dto.UpdateUserDTO) error {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found")
		}
		return err
	}

	if dto.Username != "" {
		// check if username exists where not id is the same
		userCheck, err := uc.userRepo.GetByUsername(dto.Username)
		if err != nil {
			return fmt.Errorf("failed to get user by username: %w", err)
		}
		if userCheck != nil && user.ID != userCheck.ID {
			return fmt.Errorf("username already exists")
		}
		user.Username = dto.Username
	}

	if dto.Email != "" {
		// check if email exists
		userCheck, err := uc.userRepo.GetByEmail(dto.Email)
		if err != nil {
			return fmt.Errorf("failed to get user by email: %w", err)
		}
		if userCheck != nil && user.ID != userCheck.ID {
			return fmt.Errorf("email already exists")
		}
		user.Email = dto.Email
	}

	if dto.Name != "" {
		user.Name = dto.Name
	}
	if dto.Surname != "" {
		user.Surname = dto.Surname
	}
	if dto.Role != "" {
		user.Role = dto.Role
	}
	if dto.Password != "" {
		hash, err := utils.HashPassword(dto.Password)
		if err != nil {
			return fmt.Errorf("AuthUseCase - HashPassword: %w", err)
		}
		user.Password = hash
	}

	if err := uc.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// UpdateProfile -.
func (uc *UserUseCase) UpdateProfile(id int, dto *dto.UpdateProfileDTO) error {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found")
		}
		return err
	}

	if dto.Username != "" {
		// check if username exists where not id is the same
		userCheck, err := uc.userRepo.GetByUsername(dto.Username)
		if err != nil {
			return fmt.Errorf("failed to get user by username: %w", err)
		}
		if userCheck != nil && user.ID != userCheck.ID {
			return fmt.Errorf("username already exists")
		}
		user.Username = dto.Username
	}

	if dto.Password != "" {
		hash, err := utils.HashPassword(dto.Password)
		if err != nil {
			return fmt.Errorf("AuthUseCase - HashPassword: %w", err)
		}
		user.Password = hash
	}

	if dto.Name != "" {
		user.Name = dto.Name
	}

	if dto.Surname != "" {
		user.Surname = dto.Surname
	}

	if dto.Phone != "" {
		user.Phone = dto.Phone
	}

	if err := uc.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// Delete -.
func (uc *UserUseCase) Delete(id int) error {
	user, err := uc.userRepo.GetByID(id)

	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	if err := uc.userRepo.Delete(user); err != nil {
		return err
	}

	return nil
}
