package patient

import "context"

//go:generate mockgen -destination=./mocks.go -package=domain -source=./repository.go

type Repository interface {
	CreatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	FindPatientByID(ctx context.Context, ID *uint64) (*Patient, error)
	FindAllPatientByStatus(ctx context.Context, status string) ([]*Patient, error)
	UpdatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	DeletePatient(ctx context.Context, ID *uint64) error
}

type SessionRepository interface {
	CreatePatientSession(ctx context.Context, patient *PatientSession) (*PatientSession, error)
	FindPatientByQrToken(ctx context.Context, qrToken string) (*PatientSession, error)
	UpdatePatientSession(ctx context.Context, patient *PatientSession) (*PatientSession, error)
	DeletePatientSession(ctx context.Context, ID *uint64) error
}
