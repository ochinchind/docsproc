package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineModuleRepo -.
type DisciplineModuleRepo struct {
	*postgres.Postgres
}

// NewDisciplineModuleRepo -.
func NewDisciplineModuleRepo(pg *postgres.Postgres) *DisciplineModuleRepo {
	return &DisciplineModuleRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineModuleRepo) Store(qualification *entity.DisciplineModule) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineModuleRepo) Update(disciplineModule *entity.DisciplineModule) error {
	r.Postgres.Conn.Save(&disciplineModule)

	return nil
}

// Delete -.
func (r *DisciplineModuleRepo) Delete(disciplineModule *entity.DisciplineModule) error {
	r.Postgres.Conn.Delete(&disciplineModule)

	return nil
}

// GetByID -.
func (r *DisciplineModuleRepo) GetByID(id int) (*entity.DisciplineModule, error) {
	var qualification entity.DisciplineModule
	err := r.Postgres.Conn.Where("id = ?", id).First(&qualification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves disciplineModule with pagination support
func (r *DisciplineModuleRepo) Get(ctx *gin.Context) ([]entity.DisciplineModule, int64, error) {
	var disciplineModule []entity.DisciplineModule
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

	// Count total disciplineModule
	if err := r.Postgres.Conn.Model(&entity.DisciplineModule{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated disciplineModule
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&disciplineModule).Error; err != nil {
		return nil, 0, err
	}

	return disciplineModule, total, nil
}
