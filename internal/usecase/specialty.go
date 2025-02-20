package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
)

// SpecialtyUseCase -.
type SpecialtyUseCase struct {
	specialtyRepo SpecialtyRepo
}

// NewSpecialtyUseCase -.
func NewSpecialtyUseCase(specialtyRepo SpecialtyRepo) *SpecialtyUseCase {
	return &SpecialtyUseCase{
		specialtyRepo: specialtyRepo,
	}
}

// Get -.
func (uc *SpecialtyUseCase) Get(context *gin.Context) ([]entity.Specialty, int64, error) {
	specialties, total, err := uc.specialtyRepo.Get(context)

	if err != nil {
		return nil, 0, err
	}

	return specialties, total, nil
}

// GetByID -.
func (uc *SpecialtyUseCase) GetByID(id int) (*entity.Specialty, error) {
	specialty, err := uc.specialtyRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if specialty == nil {
		return nil, errors.New("specialty not found")
	}

	return specialty, nil
}

// Update -.
func (uc *SpecialtyUseCase) Update(id int, specialty *dto.UpdateSpecialtyDTO) error {
	specialtyEntity, err := uc.specialtyRepo.GetByID(id)

	if err != nil {
		return err
	}

	if specialtyEntity == nil {
		return errors.New("specialty not found")
	}

	if specialty.Name != "" {
		specialtyEntity.Name = specialty.Name
	}

	if specialty.Code != "" {
		specialtyEntity.Code = specialty.Code
	}

	err = uc.specialtyRepo.Update(specialtyEntity)

	if err != nil {
		return err
	}

	return nil
}

// Store -.
func (uc *SpecialtyUseCase) Store(specialty *dto.StoreSpecialtyDTO) error {
	specialtyEntity := &entity.Specialty{
		Name: specialty.Name,
		Code: specialty.Code,
	}

	err := uc.specialtyRepo.Store(specialtyEntity)

	if err != nil {
		return err
	}

	return nil
}

// Delete -.
func (uc *SpecialtyUseCase) Delete(id int) error {
	specialty, err := uc.specialtyRepo.GetByID(id)

	if err != nil {
		return err
	}

	if specialty == nil {
		return errors.New("specialty not found")
	}

	err = uc.specialtyRepo.Delete(specialty)

	if err != nil {
		return err
	}

	return nil
}
