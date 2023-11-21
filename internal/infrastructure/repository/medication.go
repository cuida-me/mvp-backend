package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"gorm.io/gorm"
)

type medicationRepository struct {
	db *gorm.DB
}

func NewMedicationRepository(db *gorm.DB) *medicationRepository {
	return &medicationRepository{db: db}
}

func (r *medicationRepository) CreateMedication(ctx context.Context, medication *medication.Medication) (*medication.Medication, error) {
	if err := r.db.Create(medication).Error; err != nil {
		return nil, err
	}

	return medication, nil
}

func (r *medicationRepository) FindMedicationByID(ctx context.Context, ID *uint64) (*medication.Medication, error) {
	medication := &medication.Medication{}

	if err := r.db.Where("id = ?", ID).
		Preload("Schedules.Times").
		Preload("Type").
		First(medication).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("medication not found")
		}
		return nil, err
	}

	return medication, nil
}

func (r *medicationRepository) UpdateMedication(ctx context.Context, medication *medication.Medication) (*medication.Medication, error) {
	if err := r.db.Save(medication).Error; err != nil {
		return nil, err
	}

	return medication, nil
}

func (r *medicationRepository) DeleteMedication(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&medication.Medication{}).Error; err != nil {
		return err
	}

	return nil
}
