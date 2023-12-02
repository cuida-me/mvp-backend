package caregiver

import "context"

//go:generate mockgen -destination=./mocks.go -package=user -source=./repository.go

type Repository interface {
	CreateCaregiver(ctx context.Context, caregiver *Caregiver) (*Caregiver, error)
	FindCaregiverByID(ctx context.Context, ID *uint64) (*Caregiver, error)
	FindCaregiverByEmail(ctx context.Context, email string) (*Caregiver, error)
	FindCaregiverByUid(ctx context.Context, uid string) (*Caregiver, error)
	UpdateCaregiver(ctx context.Context, caregiver *Caregiver) (*Caregiver, error)
	DeleteCaregiver(ctx context.Context, ID *uint64) error
	FindCaregiverByPatientID(ctx context.Context, patientID *uint64) (*Caregiver, error)
}
