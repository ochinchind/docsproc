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
}

// NewDisciplineUseCase -.
func NewDisciplineUseCase(disciplineRepo DisciplineRepo, qualificationRepo QualificationRepo) *DisciplineUseCase {
	return &DisciplineUseCase{
		disciplineRepo:    disciplineRepo,
		qualificationRepo: qualificationRepo,
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

	if discipline.Competencies != "" {
		disciplineEntity.Competencies = discipline.Competencies
	}

	if discipline.Desc != "" {
		disciplineEntity.Desc = discipline.Desc
	}

	if discipline.HoursIndividual != nil {
		disciplineEntity.HoursIndividual = *discipline.HoursIndividual
	}

	if discipline.HoursInternship != nil {
		disciplineEntity.HoursInternship = *discipline.HoursInternship
	}

	if discipline.HoursPractice != nil {
		disciplineEntity.HoursPractice = *discipline.HoursPractice
	}

	if discipline.HoursSelfStudy != nil {
		disciplineEntity.HoursSelfStudy = *discipline.HoursSelfStudy
	}

	if discipline.HoursTheory != nil {
		disciplineEntity.HoursTheory = *discipline.HoursTheory
	}

	if discipline.HoursTotal != nil {
		disciplineEntity.HoursTotal = *discipline.HoursTotal
	}

	if discipline.HoursWithTeacher != nil {
		disciplineEntity.HoursWithTeacher = *discipline.HoursWithTeacher
	}

	if discipline.Lang != "" {
		disciplineEntity.Lang = discipline.Lang
	}

	if discipline.Necessities != "" {
		disciplineEntity.Necessities = discipline.Necessities
	}

	if discipline.PostRequisites != "" {
		disciplineEntity.PostRequisites = discipline.PostRequisites
	}

	if discipline.PreRequisites != "" {
		disciplineEntity.PreRequisites = discipline.PreRequisites
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

	disciplineEntity := &entity.Discipline{
		Name:             discipline.Name,
		Code:             discipline.Code,
		AssessmentType:   discipline.AssessmentType,
		Competencies:     discipline.Competencies,
		Desc:             discipline.Desc,
		HoursIndividual:  discipline.HoursIndividual,
		HoursInternship:  discipline.HoursInternship,
		HoursPractice:    discipline.HoursPractice,
		HoursSelfStudy:   discipline.HoursSelfStudy,
		HoursTheory:      discipline.HoursTheory,
		HoursTotal:       discipline.HoursTotal,
		HoursWithTeacher: discipline.HoursWithTeacher,
		Lang:             discipline.Lang,
		Necessities:      discipline.Necessities,
		PostRequisites:   discipline.PostRequisites,
		PreRequisites:    discipline.PreRequisites,
		QualificationID:  discipline.QualificationID,
		UserId:           userId,
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
