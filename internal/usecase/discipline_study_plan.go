package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineStudyPlanUseCase -.
type DisciplineStudyPlanUseCase struct {
	disciplineStudyPlanRepo DisciplineStudyPlanRepo
	disciplineRepo          DisciplineRepo
}

// NewDisciplineStudyPlanUseCase -.
func NewDisciplineStudyPlanUseCase(disciplineStudyPlanRepo DisciplineStudyPlanRepo, disciplineRepo DisciplineRepo) *DisciplineStudyPlanUseCase {
	return &DisciplineStudyPlanUseCase{
		disciplineStudyPlanRepo: disciplineStudyPlanRepo,
		disciplineRepo:          disciplineRepo,
	}
}

// Get -.
func (uc *DisciplineStudyPlanUseCase) Get(context *gin.Context) ([]entity.DisciplineStudyPlan, int64, error) {
	specialties, total, err := uc.disciplineStudyPlanRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineStudyPlanUseCase) GetByID(id int) (*entity.DisciplineStudyPlan, error) {
	disciplineStudyPlan, err := uc.disciplineStudyPlanRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if disciplineStudyPlan == nil {
		return nil, errors.New("disciplineStudyPlan not found")
	}

	return disciplineStudyPlan, nil
}

// Update -.
func (uc *DisciplineStudyPlanUseCase) Update(id int, disciplineStudyPlan *dto.UpdateDisciplineStudyPlanDTO) error {
	disciplineStudyPlanEntity, err := uc.disciplineStudyPlanRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineStudyPlanEntity == nil {
		return errors.New("disciplineStudyPlan not found")
	}

	if disciplineStudyPlan.DisciplineID != 0 {
		discipline, err := uc.disciplineRepo.GetByID(int(disciplineStudyPlan.DisciplineID))

		if err != nil {
			return err
		}

		if discipline == nil {
			return errors.New("discipline not found")
		}

		disciplineStudyPlanEntity.DisciplineID = disciplineStudyPlan.DisciplineID
	}

	if disciplineStudyPlan.Necessities != "" {
		disciplineStudyPlanEntity.Necessities = disciplineStudyPlan.Necessities
	}

	if disciplineStudyPlan.ContactInfo != "" {
		disciplineStudyPlanEntity.ContactInfo = disciplineStudyPlan.ContactInfo
	}

	if disciplineStudyPlan.PostRequisites != "" {
		disciplineStudyPlanEntity.PostRequisites = disciplineStudyPlan.PostRequisites
	}

	if disciplineStudyPlan.PreRequisites != "" {
		disciplineStudyPlanEntity.PreRequisites = disciplineStudyPlan.PreRequisites
	}

	err = uc.disciplineStudyPlanRepo.Update(disciplineStudyPlanEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineStudyPlanUseCase) Store(disciplineStudyPlan *dto.StoreDisciplineStudyPlanDTO) error {
	// check exists discipline
	discipline, err := uc.disciplineRepo.GetByID(int(disciplineStudyPlan.DisciplineID))

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	disciplineStudyPlanEntity := &entity.DisciplineStudyPlan{
		DisciplineID:   disciplineStudyPlan.DisciplineID,
		Necessities:    disciplineStudyPlan.Necessities,
		ContactInfo:    disciplineStudyPlan.ContactInfo,
		PostRequisites: disciplineStudyPlan.PostRequisites,
		PreRequisites:  disciplineStudyPlan.PreRequisites,
	}

	err = uc.disciplineStudyPlanRepo.Store(disciplineStudyPlanEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineStudyPlanUseCase) Delete(id int) error {
	disciplineStudyPlan, err := uc.disciplineStudyPlanRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineStudyPlan == nil {
		return errors.New("disciplineStudyPlan not found")
	}

	err = uc.disciplineStudyPlanRepo.Delete(disciplineStudyPlan)

	if err != nil {
		return err
	}

	return nil
}
