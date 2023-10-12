package caregiver

import "context"

//go:generate mockgen -destination=./mocks.go -package=user -source=./repository.go

type Repository interface {
	Create(ctx context.Context, user *Caregiver) (*Caregiver, error)
	FindByID(ctx context.Context, ID *int64) (*Caregiver, error)
	FindByEmail(ctx context.Context, email string) (*Caregiver, error)
	FindByUsername(ctx context.Context, username string) (*Caregiver, error)
	Update(ctx context.Context, user *Caregiver) (*Caregiver, error)
	Delete(ctx context.Context, ID *int64) error
}
