package scheduling

import "context"

type Repository interface {
	CreateScheduling(ctx context.Context, scheduling *Scheduling) (*Scheduling, error)
	FindSchedulingByID(ctx context.Context, ID *uint64) (*Scheduling, error)
	FindSchedulingByPatientAndStatus(ctx context.Context, patientID *uint64, status string) (*Scheduling, error)
	UpdateScheduling(ctx context.Context, scheduling *Scheduling) (*Scheduling, error)
	DeleteScheduling(ctx context.Context, ID *uint64) error
}
