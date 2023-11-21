package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"gorm.io/gorm"
)

type medicationScheduleRepository struct {
	db *gorm.DB
}

func NewMedicationScheduleRepository(db *gorm.DB) *medicationScheduleRepository {
	return &medicationScheduleRepository{db: db}
}

func (r *medicationScheduleRepository) CreateSchedule(ctx context.Context, schedule *medication.MedicationSchedule) (*medication.MedicationSchedule, error) {
	if err := r.db.Create(schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *medicationScheduleRepository) FindScheduleByID(ctx context.Context, ID *uint64) (*medication.MedicationSchedule, error) {
	schedule := &medication.MedicationSchedule{}

	if err := r.db.Where("id = ?", ID).First(schedule).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("medication schedule not found")
		}
		return nil, err
	}

	return schedule, nil
}

func (r *medicationScheduleRepository) UpdateSchedule(ctx context.Context, schedule *medication.MedicationSchedule) (*medication.MedicationSchedule, error) {
	if err := r.db.Save(schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *medicationScheduleRepository) DeleteSchedule(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&medication.MedicationSchedule{}).Error; err != nil {
		return err
	}

	return nil
}
