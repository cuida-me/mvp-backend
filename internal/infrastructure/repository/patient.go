package repository

import (
	"context"
	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/jinzhu/gorm"
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

func (r *patientRepository) DeletePatient(ctx context.Context, ID *int64) error {
	if err := r.db.Where("id = ?", ID).Delete(&domain.Patient{}).Error; err != nil {
		return err
	}

	return nil
}
