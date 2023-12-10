package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	"gorm.io/gorm"
)

type schedulingRepository struct {
	db *gorm.DB
}

func NewSchedulingRepository(db *gorm.DB) *schedulingRepository {
	return &schedulingRepository{db: db}
}

func (r *schedulingRepository) CreateScheduling(ctx context.Context, scheduling *scheduling.Scheduling) (*scheduling.Scheduling, error) {
	if err := r.db.Create(scheduling).Error; err != nil {
		return nil, err
	}

	return scheduling, nil
}

func (r *schedulingRepository) FindSchedulingByID(ctx context.Context, ID *uint64) (*scheduling.Scheduling, error) {
	var scheduling scheduling.Scheduling

	if err := r.db.Where("id = ?", ID).First(&scheduling).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("scheduling not found")
		}
		return nil, err
	}

	return &scheduling, nil
}

func (r *schedulingRepository) FindSchedulingByPatientAndStatus(ctx context.Context, patientID *uint64, status string) (*scheduling.Scheduling, error) {
	var scheduling scheduling.Scheduling

	if err := r.db.Where("patient_id = ? AND status = ?", patientID, status).First(&scheduling).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("scheduling not found")
		}
		return nil, err
	}

	return &scheduling, nil
}

func (r *schedulingRepository) FindAllSchedulingByMedicationID(ctx context.Context, medicationID *uint64) ([]*scheduling.Scheduling, error) {
	var schedulings []*scheduling.Scheduling

	if err := r.db.Where("medication_id = ?", medicationID).Find(&schedulings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("scheduling not found")
		}
		return nil, err
	}

	return schedulings, nil
}

func (r *schedulingRepository) UpdateScheduling(ctx context.Context, scheduling *scheduling.Scheduling) (*scheduling.Scheduling, error) {
	if err := r.db.Save(scheduling).Error; err != nil {
		return nil, err
	}

	return scheduling, nil
}

func (r *schedulingRepository) DeleteScheduling(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&scheduling.Scheduling{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *schedulingRepository) FindAllSchedulingByMedicationIDAndStatus(ctx context.Context, medicationID *uint64, status string) ([]*scheduling.Scheduling, error) {
	var schedulings []*scheduling.Scheduling

	if err := r.db.Where("medication_id = ? AND status = ?", medicationID, status).Find(&schedulings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("scheduling not found")
		}
		return nil, err
	}

	return schedulings, nil
}

func (r *schedulingRepository) FindSchedulingByMedicationIDAndDateRange(ctx context.Context, medicationID *uint64, startDate, endDate time.Time) ([]*scheduling.Scheduling, error) {
	var schedulings []*scheduling.Scheduling

	if err := r.db.Where("medication_id = ? AND medication_time BETWEEN ? AND ?", medicationID, startDate, endDate).Find(&schedulings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("scheduling not found")
		}
		return nil, err
	}

	return schedulings, nil
}
