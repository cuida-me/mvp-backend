package repository

import (
	"context"
	"fmt"
	domain "github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/jinzhu/gorm"
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

	if err := r.db.Where("id = ?", ID).First(caregiver).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("caregiver not found")
		}
		return nil, err
	}

	return caregiver, nil
}

func (r *caregiverRepository) FindCaregiverByEmail(ctx context.Context, email string) (*domain.Caregiver, error) {
	caregiver := &domain.Caregiver{}

	if err := r.db.Where("email = ?", email).First(caregiver).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
