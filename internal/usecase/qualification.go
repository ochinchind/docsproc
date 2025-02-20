package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// QualificationUseCase -.
type QualificationUseCase struct {
	qualificationRepo QualificationRepo
	specialtyRepo     SpecialtyRepo
}

// NewQualificationUseCase -.
func NewQualificationUseCase(qualificationRepo QualificationRepo, specialtyRepo SpecialtyRepo) *QualificationUseCase {
	return &QualificationUseCase{
		qualificationRepo: qualificationRepo,
		specialtyRepo:     specialtyRepo,
	}
}

// Get -.
func (uc *QualificationUseCase) Get(context *gin.Context) ([]entity.Qualification, int64, error) {
	specialties, total, err := uc.qualificationRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *QualificationUseCase) GetByID(id int) (*entity.Qualification, error) {
	qualification, err := uc.qualificationRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if qualification == nil {
		return nil, errors.New("qualification not found")
	}

	return qualification, nil
}

// Update -.
func (uc *QualificationUseCase) Update(id int, qualification *dto.UpdateQualificationDTO) error {
	qualificationEntity, err := uc.qualificationRepo.GetByID(id)

	if err != nil {
		return err
	}

	if qualificationEntity == nil {
		return errors.New("qualification not found")
	}

	if qualification.Name != "" {
		qualificationEntity.Name = qualification.Name
	}

	if qualification.Code != "" {
		qualificationEntity.Code = qualification.Code
	}

	if qualification.SpecialtyID != 0 {
		specialty, err := uc.specialtyRepo.GetByID(int(qualification.SpecialtyID))

		if err != nil {
			return err
		}

		if specialty == nil {
			return errors.New("specialty not found")
		}

		qualificationEntity.SpecialtyID = qualification.SpecialtyID
	}

	err = uc.qualificationRepo.Update(qualificationEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *QualificationUseCase) Store(qualification *dto.StoreQualificationDTO) error {
	// check exists specialty
	specialty, err := uc.specialtyRepo.GetByID(int(qualification.SpecialtyID))

	if err != nil {
		return err
	}

	if specialty == nil {
		return errors.New("specialty not found")
	}

	qualificationEntity := &entity.Qualification{
		Name:        qualification.Name,
		Code:        qualification.Code,
		SpecialtyID: qualification.SpecialtyID,
	}

	err = uc.qualificationRepo.Store(qualificationEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *QualificationUseCase) Delete(id int) error {
	qualification, err := uc.qualificationRepo.GetByID(id)

	if err != nil {
		return err
	}

	if qualification == nil {
		return errors.New("qualification not found")
	}

	err = uc.qualificationRepo.Delete(qualification)

	if err != nil {
		return err
	}

	return nil
}
