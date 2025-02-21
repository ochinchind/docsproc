package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// DisciplineModuleChapterTopicUseCase -.
type DisciplineModuleChapterTopicUseCase struct {
	disciplineModuleChapterTopicRepo DisciplineModuleChapterTopicRepo
	disciplineModuleChapterRepo      DisciplineModuleChapterRepo
}

// NewDisciplineModuleChapterTopicUseCase -.
func NewDisciplineModuleChapterTopicUseCase(disciplineModuleChapterTopicRepo DisciplineModuleChapterTopicRepo, disciplineModuleChapterRepo DisciplineModuleChapterRepo) *DisciplineModuleChapterTopicUseCase {
	return &DisciplineModuleChapterTopicUseCase{
		disciplineModuleChapterTopicRepo: disciplineModuleChapterTopicRepo,
		disciplineModuleChapterRepo:      disciplineModuleChapterRepo,
	}
}

// Get -.
func (uc *DisciplineModuleChapterTopicUseCase) Get(context *gin.Context) ([]entity.DisciplineModuleChapterTopic, int64, error) {
	specialties, total, err := uc.disciplineModuleChapterTopicRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *DisciplineModuleChapterTopicUseCase) GetByID(id int) (*entity.DisciplineModuleChapterTopic, error) {
	disciplineModuleChapterTopic, err := uc.disciplineModuleChapterTopicRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if disciplineModuleChapterTopic == nil {
		return nil, errors.New("disciplineModuleChapterTopic not found")
	}

	return disciplineModuleChapterTopic, nil
}

// Update -.
func (uc *DisciplineModuleChapterTopicUseCase) Update(id int, disciplineModuleChapterTopic *dto.UpdateDisciplineModuleChapterTopicDTO) error {
	disciplineModuleChapterTopicEntity, err := uc.disciplineModuleChapterTopicRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleChapterTopicEntity == nil {
		return errors.New("disciplineModuleChapterTopic not found")
	}

	if disciplineModuleChapterTopic.Name != "" {
		disciplineModuleChapterTopicEntity.Name = disciplineModuleChapterTopic.Name
	}

	if disciplineModuleChapterTopic.HoursIndividual != nil {
		disciplineModuleChapterTopicEntity.HoursIndividual = *disciplineModuleChapterTopic.HoursIndividual
	}

	if disciplineModuleChapterTopic.HoursInternship != nil {
		disciplineModuleChapterTopicEntity.HoursInternship = *disciplineModuleChapterTopic.HoursInternship
	}

	if disciplineModuleChapterTopic.HoursPractice != nil {
		disciplineModuleChapterTopicEntity.HoursPractice = *disciplineModuleChapterTopic.HoursPractice
	}

	if disciplineModuleChapterTopic.HoursSelfStudy != nil {
		disciplineModuleChapterTopicEntity.HoursSelfStudy = *disciplineModuleChapterTopic.HoursSelfStudy
	}

	if disciplineModuleChapterTopic.HoursTheory != nil {
		disciplineModuleChapterTopicEntity.HoursTheory = *disciplineModuleChapterTopic.HoursTheory
	}

	if disciplineModuleChapterTopic.Type != "" {
		disciplineModuleChapterTopicEntity.Type = disciplineModuleChapterTopic.Type
	}

	if disciplineModuleChapterTopic.DisciplineModuleChapterID != 0 {
		discipline, err := uc.disciplineModuleChapterRepo.GetByID(int(disciplineModuleChapterTopic.DisciplineModuleChapterID))

		if err != nil {
			return err
		}

		if discipline == nil {
			return errors.New("discipline not found")
		}

		disciplineModuleChapterTopicEntity.DisciplineModuleChapterID = disciplineModuleChapterTopic.DisciplineModuleChapterID
	}

	err = uc.disciplineModuleChapterTopicRepo.Update(disciplineModuleChapterTopicEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *DisciplineModuleChapterTopicUseCase) Store(disciplineModuleChapterTopic *dto.StoreDisciplineModuleChapterTopicDTO) error {
	// check exists discipline
	discipline, err := uc.disciplineModuleChapterRepo.GetByID(int(disciplineModuleChapterTopic.DisciplineModuleChapterID))

	if err != nil {
		return err
	}

	if discipline == nil {
		return errors.New("discipline not found")
	}

	disciplineModuleChapterTopicEntity := &entity.DisciplineModuleChapterTopic{
		Name:                      disciplineModuleChapterTopic.Name,
		DisciplineModuleChapterID: disciplineModuleChapterTopic.DisciplineModuleChapterID,
		HoursTheory:               disciplineModuleChapterTopic.HoursTheory,
		HoursPractice:             disciplineModuleChapterTopic.HoursPractice,
		HoursSelfStudy:            disciplineModuleChapterTopic.HoursSelfStudy,
		HoursInternship:           disciplineModuleChapterTopic.HoursInternship,
		HoursIndividual:           disciplineModuleChapterTopic.HoursIndividual,
		HoursWithTeacher:          disciplineModuleChapterTopic.HoursWithTeacher,
		Type:                      disciplineModuleChapterTopic.Type,
	}

	err = uc.disciplineModuleChapterTopicRepo.Store(disciplineModuleChapterTopicEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *DisciplineModuleChapterTopicUseCase) Delete(id int) error {
	disciplineModuleChapterTopic, err := uc.disciplineModuleChapterTopicRepo.GetByID(id)

	if err != nil {
		return err
	}

	if disciplineModuleChapterTopic == nil {
		return errors.New("disciplineModuleChapterTopic not found")
	}

	err = uc.disciplineModuleChapterTopicRepo.Delete(disciplineModuleChapterTopic)

	if err != nil {
		return err
	}

	return nil
}
