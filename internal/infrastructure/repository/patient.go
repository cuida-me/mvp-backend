package repository

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *patientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) CreatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error) {
	if err := r.db.Create(patient).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *patientRepository) FindPatientByID(ctx context.Context, ID *uint64) (*domain.Patient, error) {
	patient := &domain.Patient{}

	if err := r.db.Where("id = ?", ID).First(patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("patient not found")
		}
		return nil, err
	}

	return patient, nil
}

func (r *patientRepository) UpdatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error) {
	if err := r.db.Save(patient).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *patientRepository) DeletePatient(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&domain.Patient{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *patientRepository) FindAllPatientByStatus(ctx context.Context, status string) ([]*domain.Patient, error) {
	var patients []*domain.Patient

	if err := r.db.Where("status = ?", status).Find(&patients).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("patient not found")
		}
		return nil, err
	}

	return patients, nil
}
