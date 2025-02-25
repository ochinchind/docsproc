package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineStudyPlanRepo -.
type DisciplineStudyPlanRepo struct {
	*postgres.Postgres
}

// NewDisciplineStudyPlanRepo -.
func NewDisciplineStudyPlanRepo(pg *postgres.Postgres) *DisciplineStudyPlanRepo {
	return &DisciplineStudyPlanRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineStudyPlanRepo) Store(qualification *entity.DisciplineStudyPlan) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineStudyPlanRepo) Update(disciplineStudyPlan *entity.DisciplineStudyPlan) error {
	r.Postgres.Conn.Save(&disciplineStudyPlan)

	return nil
}

// Delete -.
func (r *DisciplineStudyPlanRepo) Delete(disciplineStudyPlan *entity.DisciplineStudyPlan) error {
	r.Postgres.Conn.Delete(&disciplineStudyPlan)

	return nil
}

// GetByID -.
func (r *DisciplineStudyPlanRepo) GetByID(id int) (*entity.DisciplineStudyPlan, error) {
	var qualification entity.DisciplineStudyPlan
	err := r.Postgres.Conn.
		Preload("DisciplineStudyPlanTopics").
		Where("id = ?", id).
		First(&qualification).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves disciplineStudyPlan with pagination support
func (r *DisciplineStudyPlanRepo) Get(ctx *gin.Context) ([]entity.DisciplineStudyPlan, int64, error) {
	var disciplineStudyPlan []entity.DisciplineStudyPlan
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

	// Count total disciplineStudyPlan
	if err := r.Postgres.Conn.Model(&entity.DisciplineStudyPlan{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated disciplineStudyPlan
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&disciplineStudyPlan).Error; err != nil {
		return nil, 0, err
	}

	return disciplineStudyPlan, total, nil
}
