package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineRepo -.
type DisciplineRepo struct {
	*postgres.Postgres
}

// NewDisciplineRepo -.
func NewDisciplineRepo(pg *postgres.Postgres) *DisciplineRepo {
	return &DisciplineRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineRepo) Store(qualification *entity.Discipline) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineRepo) Update(discipline *entity.Discipline) error {
	r.Postgres.Conn.Save(&discipline)

	return nil
}

// Delete -.
func (r *DisciplineRepo) Delete(discipline *entity.Discipline) error {
	r.Postgres.Conn.Delete(&discipline)

	return nil
}

// GetByID -.
func (r *DisciplineRepo) GetByID(id int) (*entity.Discipline, error) {
	var qualification entity.Discipline
	err := r.Postgres.Conn.Where("id = ?", id).First(&qualification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves discipline with pagination support
func (r *DisciplineRepo) Get(ctx *gin.Context) ([]entity.Discipline, int64, error) {
	var discipline []entity.Discipline
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

	// Count total discipline
	if err := r.Postgres.Conn.Model(&entity.Discipline{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated discipline
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&discipline).Error; err != nil {
		return nil, 0, err
	}

	return discipline, total, nil
}
