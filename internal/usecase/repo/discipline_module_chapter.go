package repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/pkg/postgres"
	"gorm.io/gorm"
	"strconv"
)

// DisciplineModuleChapterRepo -.
type DisciplineModuleChapterRepo struct {
	*postgres.Postgres
}

// NewDisciplineModuleChapterRepo -.
func NewDisciplineModuleChapterRepo(pg *postgres.Postgres) *DisciplineModuleChapterRepo {
	return &DisciplineModuleChapterRepo{
		Postgres: pg,
	}
}

// Store -.
func (r *DisciplineModuleChapterRepo) Store(qualification *entity.DisciplineModuleChapter) error {
	err := r.Postgres.Conn.Create(&qualification).Error

	return err
}

// Update -.
func (r *DisciplineModuleChapterRepo) Update(disciplineModuleChapter *entity.DisciplineModuleChapter) error {
	r.Postgres.Conn.Save(&disciplineModuleChapter)

	return nil
}

// Delete -.
func (r *DisciplineModuleChapterRepo) Delete(disciplineModuleChapter *entity.DisciplineModuleChapter) error {
	r.Postgres.Conn.Delete(&disciplineModuleChapter)

	return nil
}

// GetByID -.
func (r *DisciplineModuleChapterRepo) GetByID(id int) (*entity.DisciplineModuleChapter, error) {
	var qualification entity.DisciplineModuleChapter
	err := r.Postgres.Conn.
		Preload("DisciplineModuleChapterTopics").
		Where("id = ?", id).
		First(&qualification).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &qualification, err
}

// Get - Retrieves disciplineModuleChapter with pagination support
func (r *DisciplineModuleChapterRepo) Get(ctx *gin.Context) ([]entity.DisciplineModuleChapter, int64, error) {
	var disciplineModuleChapter []entity.DisciplineModuleChapter
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

	// Count total disciplineModuleChapter
	if err := r.Postgres.Conn.Model(&entity.DisciplineModuleChapter{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated disciplineModuleChapter
	if err := r.Postgres.Conn.Limit(limit).Offset(offset).Find(&disciplineModuleChapter).Error; err != nil {
		return nil, 0, err
	}

	return disciplineModuleChapter, total, nil
}
