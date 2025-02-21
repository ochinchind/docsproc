package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineModuleChapterTopicRepo -.
type DisciplineModuleChapterTopicRepo struct {
	*postgres.Postgres
}

// NewDisciplineModuleChapterTopicRepo -.
func NewDisciplineModuleChapterTopicRepo(pg *postgres.Postgres) *DisciplineModuleChapterTopicRepo {
	return &DisciplineModuleChapterTopicRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineModuleChapterTopicRepo) Store(qualification *entity.DisciplineModuleChapterTopic) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineModuleChapterTopicRepo) Update(disciplineModuleChapterTopic *entity.DisciplineModuleChapterTopic) error {
	r.Postgres.Conn.Save(&disciplineModuleChapterTopic)

	return nil
}

// Delete -.
func (r *DisciplineModuleChapterTopicRepo) Delete(disciplineModuleChapterTopic *entity.DisciplineModuleChapterTopic) error {
	r.Postgres.Conn.Delete(&disciplineModuleChapterTopic)

	return nil
}

// GetByID -.
func (r *DisciplineModuleChapterTopicRepo) GetByID(id int) (*entity.DisciplineModuleChapterTopic, error) {
	var qualification entity.DisciplineModuleChapterTopic
	err := r.Postgres.Conn.Where("id = ?", id).First(&qualification).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves disciplineModuleChapterTopic with pagination support
func (r *DisciplineModuleChapterTopicRepo) Get(ctx *gin.Context) ([]entity.DisciplineModuleChapterTopic, int64, error) {
	var disciplineModuleChapterTopic []entity.DisciplineModuleChapterTopic
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

	// Count total disciplineModuleChapterTopic
	if err := r.Postgres.Conn.Model(&entity.DisciplineModuleChapterTopic{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated disciplineModuleChapterTopic
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&disciplineModuleChapterTopic).Error; err != nil {
		return nil, 0, err
	}

	return disciplineModuleChapterTopic, total, nil
}
