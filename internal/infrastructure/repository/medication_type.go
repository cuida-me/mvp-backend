package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"gorm.io/gorm"
)

type medicationTypeRepository struct {
	db *gorm.DB
}

func NewMedicationTypeRepository(db *gorm.DB) *medicationTypeRepository {
	return &medicationTypeRepository{db: db}
}

func (r *medicationTypeRepository) FindAllTypes(ctx context.Context) ([]*medication.MedicationType, error) {
	var types []*medication.MedicationType

	if err := r.db.Find(&types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func (r *medicationTypeRepository) FindTypeByID(ctx context.Context, ID *uint64) (*medication.MedicationType, error) {
	medicationType := &medication.MedicationType{}

	if err := r.db.Where("id = ?", ID).First(medicationType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("medication type not found")
		}
		return nil, err
	}

	return medicationType, nil
}
