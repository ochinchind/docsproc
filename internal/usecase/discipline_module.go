package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineModuleUseCase -.
type DisciplineModuleUseCase struct {
	disciplineModuleRepo DisciplineModuleRepo
	disciplineRepo       DisciplineRepo
}

// NewDisciplineModuleUseCase -.
func NewDisciplineModuleUseCase(disciplineModuleRepo DisciplineModuleRepo, disciplineRepo DisciplineRepo) *DisciplineModuleUseCase {
	return &DisciplineModuleUseCase{
		disciplineModuleRepo: disciplineModuleRepo,
		disciplineRepo:       disciplineRepo,
	}
}

// Get -.
func (uc *DisciplineModuleUseCase) Get(context *gin.Context) ([]entity.DisciplineModule, int64, error) {
	specialties, total, err := uc.disciplineModuleRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineModuleUseCase) GetByID(id int) (*entity.DisciplineModule, error) {
	disciplineModule, err := uc.disciplineModuleRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if disciplineModule == nil {
		return nil, errors.New("disciplineModule not found")
	}

	return disciplineModule, nil
}

// Update -.
func (uc *DisciplineModuleUseCase) Update(id int, disciplineModule *dto.UpdateDisciplineModuleDTO) error {
	disciplineModuleEntity, err := uc.disciplineModuleRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleEntity == nil {
		return errors.New("disciplineModule not found")
	}

	if disciplineModule.Name != "" {
		disciplineModuleEntity.Name = disciplineModule.Name
	}

	if disciplineModule.DisciplineID != 0 {
		discipline, err := uc.disciplineRepo.GetByID(int(disciplineModule.DisciplineID))

		if err != nil {
			return err
		}

		if discipline == nil {
			return errors.New("discipline not found")
		}

		disciplineModuleEntity.DisciplineID = disciplineModule.DisciplineID
	}

	err = uc.disciplineModuleRepo.Update(disciplineModuleEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineModuleUseCase) Store(disciplineModule *dto.StoreDisciplineModuleDTO) error {
	// check exists discipline
	discipline, err := uc.disciplineRepo.GetByID(int(disciplineModule.DisciplineID))

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	disciplineModuleEntity := &entity.DisciplineModule{
		Name:         disciplineModule.Name,
		DisciplineID: disciplineModule.DisciplineID,
	}

	err = uc.disciplineModuleRepo.Store(disciplineModuleEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineModuleUseCase) Delete(id int) error {
	disciplineModule, err := uc.disciplineModuleRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModule == nil {
		return errors.New("disciplineModule not found")
	}

	err = uc.disciplineModuleRepo.Delete(disciplineModule)

	if err != nil {
		return err
	}

	return nil
}
