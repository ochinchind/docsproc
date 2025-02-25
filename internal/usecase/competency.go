package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// CompetencyUseCase -.
type CompetencyUseCase struct {
	competencyRepo CompetencyRepo
	specialtyRepo  SpecialtyRepo
}

// NewCompetencyUseCase -.
func NewCompetencyUseCase(competencyRepo CompetencyRepo, specialtyRepo SpecialtyRepo) *CompetencyUseCase {
	return &CompetencyUseCase{
		competencyRepo: competencyRepo,
		specialtyRepo:  specialtyRepo,
	}
}

// Get -.
func (uc *CompetencyUseCase) Get(context *gin.Context) ([]entity.Competency, int64, error) {
	specialties, total, err := uc.competencyRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *CompetencyUseCase) GetByID(id int) (*entity.Competency, error) {
	competency, err := uc.competencyRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if competency == nil {
		return nil, errors.New("competency not found")
	}

	return competency, nil
}

// Update -.
func (uc *CompetencyUseCase) Update(id int, competency *dto.UpdateCompetencyDTO) error {
	competencyEntity, err := uc.competencyRepo.GetByID(id)

	if err != nil {
		return err
	}

	if competencyEntity == nil {
		return errors.New("competency not found")
	}

	if competency.Name != "" {
		competencyEntity.Name = competency.Name
	}

	err = uc.competencyRepo.Update(competencyEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *CompetencyUseCase) Store(competency *dto.StoreCompetencyDTO) error {
	competencyEntity := &entity.Competency{
		Name: competency.Name,
	}

	err := uc.competencyRepo.Store(competencyEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *CompetencyUseCase) Delete(id int) error {
	competency, err := uc.competencyRepo.GetByID(id)

	if err != nil {
		return err
	}

	if competency == nil {
		return errors.New("competency not found")
	}

	err = uc.competencyRepo.Delete(competency)

	if err != nil {
		return err
	}

	return nil
}
