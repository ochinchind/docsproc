package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineModuleTopicRepo -.
type DisciplineModuleTopicRepo struct {
	*postgres.Postgres
}

// NewDisciplineModuleTopicRepo -.
func NewDisciplineModuleTopicRepo(pg *postgres.Postgres) *DisciplineModuleTopicRepo {
	return &DisciplineModuleTopicRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineModuleTopicRepo) Store(qualification *entity.DisciplineModuleTopic) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineModuleTopicRepo) Update(disciplineModuleTopic *entity.DisciplineModuleTopic) error {
	r.Postgres.Conn.Save(&disciplineModuleTopic)

	return nil
}

// Delete -.
func (r *DisciplineModuleTopicRepo) Delete(disciplineModuleTopic *entity.DisciplineModuleTopic) error {
	r.Postgres.Conn.Delete(&disciplineModuleTopic)

	return nil
}

// GetByID -.
func (r *DisciplineModuleTopicRepo) GetByID(id int) (*entity.DisciplineModuleTopic, error) {
	var qualification entity.DisciplineModuleTopic
	err := r.Postgres.Conn.Where("id = ?", id).First(&qualification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves disciplineModuleTopic with pagination support
func (r *DisciplineModuleTopicRepo) Get(ctx *gin.Context) ([]entity.DisciplineModuleTopic, int64, error) {
	var disciplineModuleTopic []entity.DisciplineModuleTopic
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

	// Count total disciplineModuleTopic
	if err := r.Postgres.Conn.Model(&entity.DisciplineModuleTopic{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated disciplineModuleTopic
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&disciplineModuleTopic).Error; err != nil {
		return nil, 0, err
	}

	return disciplineModuleTopic, total, nil
}
