package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineModuleChapterUseCase -.
type DisciplineModuleChapterUseCase struct {
	disciplineModuleChapterRepo DisciplineModuleChapterRepo
	disciplineModuleRepo        DisciplineModuleRepo
}

// NewDisciplineModuleChapterUseCase -.
func NewDisciplineModuleChapterUseCase(disciplineModuleChapterRepo DisciplineModuleChapterRepo, disciplineModuleRepo DisciplineModuleRepo) *DisciplineModuleChapterUseCase {
	return &DisciplineModuleChapterUseCase{
		disciplineModuleChapterRepo: disciplineModuleChapterRepo,
		disciplineModuleRepo:        disciplineModuleRepo,
	}
}

// Get -.
func (uc *DisciplineModuleChapterUseCase) Get(context *gin.Context) ([]entity.DisciplineModuleChapter, int64, error) {
	specialties, total, err := uc.disciplineModuleChapterRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineModuleChapterUseCase) GetByID(id int) (*entity.DisciplineModuleChapter, error) {
	disciplineModuleChapter, err := uc.disciplineModuleChapterRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if disciplineModuleChapter == nil {
		return nil, errors.New("disciplineModuleChapter not found")
	}

	return disciplineModuleChapter, nil
}

// Update -.
func (uc *DisciplineModuleChapterUseCase) Update(id int, disciplineModuleChapter *dto.UpdateDisciplineModuleChapterDTO) error {
	disciplineModuleChapterEntity, err := uc.disciplineModuleChapterRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleChapterEntity == nil {
		return errors.New("disciplineModuleChapter not found")
	}

	if disciplineModuleChapter.Name != "" {
		disciplineModuleChapterEntity.Name = disciplineModuleChapter.Name
	}

	if disciplineModuleChapter.DisciplineModuleID != 0 {
		discipline, err := uc.disciplineModuleRepo.GetByID(int(disciplineModuleChapter.DisciplineModuleID))

		if err != nil {
			return err
		}

		if discipline == nil {
			return errors.New("discipline not found")
		}

		disciplineModuleChapterEntity.DisciplineModuleID = disciplineModuleChapter.DisciplineModuleID
	}

	err = uc.disciplineModuleChapterRepo.Update(disciplineModuleChapterEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineModuleChapterUseCase) Store(disciplineModuleChapter *dto.StoreDisciplineModuleChapterDTO) error {
	// check exists discipline
	discipline, err := uc.disciplineModuleRepo.GetByID(int(disciplineModuleChapter.DisciplineModuleID))

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	disciplineModuleChapterEntity := &entity.DisciplineModuleChapter{
		Name:               disciplineModuleChapter.Name,
		DisciplineModuleID: disciplineModuleChapter.DisciplineModuleID,
	}

	err = uc.disciplineModuleChapterRepo.Store(disciplineModuleChapterEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineModuleChapterUseCase) Delete(id int) error {
	disciplineModuleChapter, err := uc.disciplineModuleChapterRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleChapter == nil {
		return errors.New("disciplineModuleChapter not found")
	}

	err = uc.disciplineModuleChapterRepo.Delete(disciplineModuleChapter)

	if err != nil {
		return err
	}

	return nil
}
