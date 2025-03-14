package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{
		Postgres: pg,
	}
}

// GetByUsernameOrEmail -.
func (r *UserRepo) GetByUsernameOrEmail(username, email string) (*entity.User, error) {
	var user entity.User
	err := r.Postgres.Conn.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByUsername -.
func (r *UserRepo) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.Postgres.Conn.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByEmail -.
func (r *UserRepo) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.Postgres.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// GetByID -.
func (r *UserRepo) GetByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.Postgres.Conn.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Create -.
func (r *UserRepo) Create(user *entity.User) error {
	r.Postgres.Conn.Create(&user)

	return nil
}

// Update -.
func (r *UserRepo) Update(user *entity.User) error {
	r.Postgres.Conn.Save(&user)

	return nil
}

// Delete -.
func (r *UserRepo) Delete(user *entity.User) error {
	r.Postgres.Conn.Delete(&user)

	return nil
}

// Get - Retrieves users with pagination support
func (r *UserRepo) Get(ctx *gin.Context) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64

	// Get page & limit from query parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Count total users
	if err := r.Postgres.Conn.Model(&entity.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated users
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
