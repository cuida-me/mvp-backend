package repository

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"gorm.io/gorm"
)

type caregiverRepository struct {
	db *gorm.DB
}

func NewCaregiverRepository(db *gorm.DB) *caregiverRepository {
	return &caregiverRepository{db: db}
}

func (r *caregiverRepository) CreateCaregiver(ctx context.Context, caregiver *domain.Caregiver) (*domain.Caregiver, error) {
	if err := r.db.Create(caregiver).Error; err != nil {
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) FindCaregiverByID(ctx context.Context, ID *uint64) (*domain.Caregiver, error) {
	caregiver := &domain.Caregiver{}

	if err := r.db.Where("id = ?", ID).Preload("Patient").First(caregiver).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("caregiver not found")
		}
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) FindCaregiverByEmail(ctx context.Context, email string) (*domain.Caregiver, error) {
	caregiver := &domain.Caregiver{}

	if err := r.db.Where("email = ?", email).First(caregiver).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("caregiver not found")
		}
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) UpdateCaregiver(ctx context.Context, caregiver *domain.Caregiver) (*domain.Caregiver, error) {
	if err := r.db.Save(caregiver).Error; err != nil {
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) DeleteCaregiver(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&domain.Caregiver{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *caregiverRepository) FindCaregiverByPatientID(ctx context.Context, patientID *uint64) (*domain.Caregiver, error) {
	caregiver := &domain.Caregiver{}

	if err := r.db.Where("patient_id = ?", patientID).Preload("Patient").First(caregiver).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("caregiver not found")
		}
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) FindCaregiverByUid(ctx context.Context, uid string) (*domain.Caregiver, error) {
	caregiver := &domain.Caregiver{}

	if err := r.db.Where("uid = ?", uid).Preload("Patient").First(caregiver).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("caregiver not found")
		}
		return nil, err
	}

	return caregiver, nil
}
