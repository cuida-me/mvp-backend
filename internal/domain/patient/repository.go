package patient

import "context"

//go:generate mockgen -destination=./mocks.go -package=domain -source=./repository.go

type Repository interface {
	CreatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	FindPatientByID(ctx context.Context, ID *uint64) (*Patient, error)
	UpdatePatient(ctx context.Context, patient *Patient) (*Patient, error)
	DeletePatient(ctx context.Context, ID *int64) error
}
