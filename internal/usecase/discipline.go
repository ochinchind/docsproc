package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineUseCase -.
type DisciplineUseCase struct {
	disciplineRepo    DisciplineRepo
	qualificationRepo QualificationRepo
	competencyRepo    CompetencyRepo
}

// NewDisciplineUseCase -.
func NewDisciplineUseCase(disciplineRepo DisciplineRepo, qualificationRepo QualificationRepo, competencyRepo CompetencyRepo) *DisciplineUseCase {
	return &DisciplineUseCase{
		disciplineRepo:    disciplineRepo,
		qualificationRepo: qualificationRepo,
		competencyRepo:    competencyRepo,
	}
}

// Get -.
func (uc *DisciplineUseCase) Get(context *gin.Context) ([]entity.Discipline, int64, error) {
	specialties, total, err := uc.disciplineRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineUseCase) GetByID(id int) (*entity.Discipline, error) {
	discipline, err := uc.disciplineRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if discipline == nil {
		return nil, errors.New("discipline not found")
	}

	return discipline, nil
}

// Update -.
func (uc *DisciplineUseCase) Update(id int, discipline *dto.UpdateDisciplineDTO) error {
	disciplineEntity, err := uc.disciplineRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineEntity == nil {
		return errors.New("discipline not found")
	}

	if discipline.Name != "" {
		disciplineEntity.Name = discipline.Name
	}

	if discipline.Code != "" {
		disciplineEntity.Code = discipline.Code
	}

	if discipline.AssessmentType != "" {
		disciplineEntity.AssessmentType = discipline.AssessmentType
	}

	if discipline.CompetencyID != 0 {
		competency, err := uc.competencyRepo.GetByID(int(discipline.CompetencyID))

		if err != nil {
			return err
		}

		if competency == nil {
			return errors.New("competency not found")
		}

		disciplineEntity.CompetencyID = discipline.CompetencyID
	}

	if discipline.Desc != "" {
		disciplineEntity.Desc = discipline.Desc
	}

	if discipline.HoursTotal != nil {
		disciplineEntity.HoursTotal = *discipline.HoursTotal
	}

	if discipline.CreaditsCount != nil {
		disciplineEntity.CreaditsCount = *discipline.CreaditsCount
	}

	if discipline.EducationForm != "" {
		disciplineEntity.EducationForm = discipline.EducationForm
	}

	if discipline.EducationBase != "" {
		disciplineEntity.EducationBase = discipline.EducationBase
	}

	if discipline.Lang != "" {
		disciplineEntity.Lang = discipline.Lang
	}

	if discipline.QualificationID != 0 {
		qualification, err := uc.qualificationRepo.GetByID(int(discipline.QualificationID))

		if err != nil {
			return err
		}

		if qualification == nil {
			return errors.New("qualification not found")
		}

		disciplineEntity.QualificationID = discipline.QualificationID
	}

	err = uc.disciplineRepo.Update(disciplineEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineUseCase) Store(discipline *dto.StoreDisciplineDTO, userId uint) error {
	// check exists qualification
	qualification, err := uc.qualificationRepo.GetByID(int(discipline.QualificationID))

	if err != nil {
		return err
	}

	if qualification == nil {
		return errors.New("qualification not found")
	}

	competency, err := uc.competencyRepo.GetByID(int(discipline.CompetencyID))

	if err != nil {
		return err
	}

	if competency == nil {
		return errors.New("competency not found")
	}

	disciplineEntity := &entity.Discipline{
		Name:            discipline.Name,
		Code:            discipline.Code,
		AssessmentType:  discipline.AssessmentType,
		CompetencyID:    discipline.CompetencyID,
		Desc:            discipline.Desc,
		CreaditsCount:   discipline.CreaditsCount,
		EducationForm:   discipline.EducationForm,
		EducationBase:   discipline.EducationBase,
		HoursTotal:      discipline.HoursTotal,
		Lang:            discipline.Lang,
		QualificationID: discipline.QualificationID,
	}

	err = uc.disciplineRepo.Store(disciplineEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineUseCase) Delete(id int) error {
	discipline, err := uc.disciplineRepo.GetByID(id)

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	err = uc.disciplineRepo.Delete(discipline)

	if err != nil {
		return err
	}

	return nil
}
