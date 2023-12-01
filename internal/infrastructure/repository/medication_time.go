package repository

import (
	"context"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"gorm.io/gorm"
)

type medicationTimeRepository struct {
	db *gorm.DB
}

func NewMedicationTimeRepository(db *gorm.DB) *medicationTimeRepository {
	return &medicationTimeRepository{db: db}
}

func (r *medicationTimeRepository) DeleteTime(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&medication.MedicationTime{}).Error; err != nil {
		return err
	}

	return nil
}
