package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// SpecialtyRepo -.
type SpecialtyRepo struct {
	*postgres.Postgres
}

// NewSpecialtyRepo -.
func NewSpecialtyRepo(pg *postgres.Postgres) *SpecialtyRepo {
	return &SpecialtyRepo{
		Postgres: pg,
	}
}

// Create -.
func (r *SpecialtyRepo) Create(specialties *entity.Specialty) error {
	r.Postgres.Conn.Create(&specialties)

	return nil
}

// Update -.
func (r *SpecialtyRepo) Update(specialties *entity.Specialty) error {
	r.Postgres.Conn.Save(&specialties)

	return nil
}

// Delete -.
func (r *SpecialtyRepo) Delete(specialties *entity.Specialty) error {
	r.Postgres.Conn.Delete(&specialties)

	return nil
}

// GetByID -.
func (r *SpecialtyRepo) GetByID(id int) (*entity.Specialty, error) {
	var specialty entity.Specialty
	err := r.Postgres.Conn.Where("id = ?", id).First(&specialty).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &specialty, err
}

// Get - Retrieves specialties with pagination support
func (r *SpecialtyRepo) Get(ctx *gin.Context) ([]entity.Specialty, int64, error) {
	var specialties []entity.Specialty
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

	// Count total specialties
	if err := r.Postgres.Conn.Model(&entity.Specialty{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated specialties
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&specialties).Error; err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}
