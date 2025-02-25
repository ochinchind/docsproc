package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// CompetencyRepo -.
type CompetencyRepo struct {
	*postgres.Postgres
}

// NewCompetencyRepo -.
func NewCompetencyRepo(pg *postgres.Postgres) *CompetencyRepo {
	return &CompetencyRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *CompetencyRepo) Store(competency *entity.Competency) error {
	err := r.Postgres.Conn.Create(&competency).Error

	return err
}

// Update -.
func (r *CompetencyRepo) Update(competencies *entity.Competency) error {
	r.Postgres.Conn.Save(&competencies)

	return nil
}

// Delete -.
func (r *CompetencyRepo) Delete(competencies *entity.Competency) error {
	r.Postgres.Conn.Delete(&competencies)

	return nil
}

// GetByID -.
func (r *CompetencyRepo) GetByID(id int) (*entity.Competency, error) {
	var competency entity.Competency
	err := r.Postgres.Conn.Where("id = ?", id).First(&competency).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &competency, err
}

// Get - Retrieves competencies with pagination support
func (r *CompetencyRepo) Get(ctx *gin.Context) ([]entity.Competency, int64, error) {
	var competencies []entity.Competency
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

	// Count total competencies
	if err := r.Postgres.Conn.Model(&entity.Competency{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated competencies
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&competencies).Error; err != nil {
		return nil, 0, err
	}

	return competencies, total, nil
}
