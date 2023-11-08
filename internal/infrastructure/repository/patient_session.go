package repository

import (
	"context"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"gorm.io/gorm"
)

type patientSessionRepository struct {
	db *gorm.DB
}

func NewPatientSessionRepository(db *gorm.DB) *patientSessionRepository {
	return &patientSessionRepository{db: db}
}

func (r patientSessionRepository) CreatePatientSession(ctx context.Context, patient *patient.PatientSession) (*patient.PatientSession, error) {
	if err := r.db.Create(patient).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

func (r patientSessionRepository) FindPatientByQrToken(ctx context.Context, qrToken string) (*patient.PatientSession, error) {
	patientSession := &patient.PatientSession{}

	if err := r.db.Where("qr_token = ?", qrToken).First(patientSession).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return patientSession, nil
}

func (r patientSessionRepository) UpdatePatientSession(ctx context.Context, patient *patient.PatientSession) (*patient.PatientSession, error) {
	if err := r.db.Save(patient).Error; err != nil {
		return nil, err
	}

	return patient, nil
}

func (r patientSessionRepository) DeletePatientSession(ctx context.Context, ID *uint64) error {
	if err := r.db.Where("id = ?", ID).Delete(&patient.PatientSession{}).Error; err != nil {
		return err
	}

	return nil
}
