package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineModuleTopicUseCase -.
type DisciplineModuleTopicUseCase struct {
	disciplineModuleTopicRepo DisciplineModuleTopicRepo
	disciplineModuleRepo      DisciplineModuleRepo
}

// NewDisciplineModuleTopicUseCase -.
func NewDisciplineModuleTopicUseCase(disciplineModuleTopicRepo DisciplineModuleTopicRepo, disciplineModuleRepo DisciplineModuleRepo) *DisciplineModuleTopicUseCase {
	return &DisciplineModuleTopicUseCase{
		disciplineModuleTopicRepo: disciplineModuleTopicRepo,
		disciplineModuleRepo:      disciplineModuleRepo,
	}
}

// Get -.
func (uc *DisciplineModuleTopicUseCase) Get(context *gin.Context) ([]entity.DisciplineModuleTopic, int64, error) {
	specialties, total, err := uc.disciplineModuleTopicRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineModuleTopicUseCase) GetByID(id int) (*entity.DisciplineModuleTopic, error) {
	disciplineModuleTopic, err := uc.disciplineModuleTopicRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if disciplineModuleTopic == nil {
		return nil, errors.New("disciplineModuleTopic not found")
	}

	return disciplineModuleTopic, nil
}

// Update -.
func (uc *DisciplineModuleTopicUseCase) Update(id int, disciplineModuleTopic *dto.UpdateDisciplineModuleTopicDTO) error {
	disciplineModuleTopicEntity, err := uc.disciplineModuleTopicRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleTopicEntity == nil {
		return errors.New("disciplineModuleTopic not found")
	}

	if disciplineModuleTopic.Name != "" {
		disciplineModuleTopicEntity.Name = disciplineModuleTopic.Name
	}

	if disciplineModuleTopic.HoursIndividual != 0 {
		disciplineModuleTopicEntity.HoursIndividual = disciplineModuleTopic.HoursIndividual
	}

	if disciplineModuleTopic.HoursInternship != 0 {
		disciplineModuleTopicEntity.HoursInternship = disciplineModuleTopic.HoursInternship
	}

	if disciplineModuleTopic.HoursPractice != 0 {
		disciplineModuleTopicEntity.HoursPractice = disciplineModuleTopic.HoursPractice
	}

	if disciplineModuleTopic.HoursSelfStudy != 0 {
		disciplineModuleTopicEntity.HoursSelfStudy = disciplineModuleTopic.HoursSelfStudy
	}

	if disciplineModuleTopic.HoursTheory != 0 {
		disciplineModuleTopicEntity.HoursTheory = disciplineModuleTopic.HoursTheory
	}

	if disciplineModuleTopic.Type != "" {
		disciplineModuleTopicEntity.Type = disciplineModuleTopic.Type
	}

	if disciplineModuleTopic.DisciplineModuleID != 0 {
		discipline, err := uc.disciplineModuleRepo.GetByID(int(disciplineModuleTopic.DisciplineModuleID))

		if err != nil {
			return err
		}

		if discipline == nil {
			return errors.New("discipline not found")
		}

		disciplineModuleTopicEntity.DisciplineModuleID = disciplineModuleTopic.DisciplineModuleID
	}

	err = uc.disciplineModuleTopicRepo.Update(disciplineModuleTopicEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineModuleTopicUseCase) Store(disciplineModuleTopic *dto.StoreDisciplineModuleTopicDTO) error {
	// check exists discipline
	discipline, err := uc.disciplineModuleRepo.GetByID(int(disciplineModuleTopic.DisciplineModuleID))

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	disciplineModuleTopicEntity := &entity.DisciplineModuleTopic{
		Name:               disciplineModuleTopic.Name,
		DisciplineModuleID: disciplineModuleTopic.DisciplineModuleID,
		HoursTheory:        disciplineModuleTopic.HoursTheory,
		HoursPractice:      disciplineModuleTopic.HoursPractice,
		HoursSelfStudy:     disciplineModuleTopic.HoursSelfStudy,
		HoursInternship:    disciplineModuleTopic.HoursInternship,
		HoursIndividual:    disciplineModuleTopic.HoursIndividual,
		HoursWithTeacher:   disciplineModuleTopic.HoursWithTeacher,
		Type:               disciplineModuleTopic.Type,
	}

	err = uc.disciplineModuleTopicRepo.Store(disciplineModuleTopicEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineModuleTopicUseCase) Delete(id int) error {
	disciplineModuleTopic, err := uc.disciplineModuleTopicRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleTopic == nil {
		return errors.New("disciplineModuleTopic not found")
	}

	err = uc.disciplineModuleTopicRepo.Delete(disciplineModuleTopic)

	if err != nil {
		return err
	}

	return nil
}
