package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// QualificationRepo -.
type QualificationRepo struct {
	*postgres.Postgres
}

// NewQualificationRepo -.
func NewQualificationRepo(pg *postgres.Postgres) *QualificationRepo {
	return &QualificationRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *QualificationRepo) Store(qualification *entity.Qualification) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *QualificationRepo) Update(qualifications *entity.Qualification) error {
	r.Postgres.Conn.Save(&qualifications)

	return nil
}

// Delete -.
func (r *QualificationRepo) Delete(qualifications *entity.Qualification) error {
	r.Postgres.Conn.Delete(&qualifications)

	return nil
}

// GetByID -.
func (r *QualificationRepo) GetByID(id int) (*entity.Qualification, error) {
	var qualification entity.Qualification
	err := r.Postgres.Conn.Where("id = ?", id).First(&qualification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves qualifications with pagination support
func (r *QualificationRepo) Get(ctx *gin.Context) ([]entity.Qualification, int64, error) {
	var qualifications []entity.Qualification
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

	// Count total qualifications
	if err := r.Postgres.Conn.Model(&entity.Qualification{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated qualifications
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&qualifications).Error; err != nil {
		return nil, 0, err
	}

	return qualifications, total, nil
}
