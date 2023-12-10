package scheduling

import (
	"context"
	"time"
)

type Repository interface {
	CreateScheduling(ctx context.Context, scheduling *Scheduling) (*Scheduling, error)
	FindSchedulingByID(ctx context.Context, ID *uint64) (*Scheduling, error)
	FindSchedulingByPatientAndStatus(ctx context.Context, patientID *uint64, status string) (*Scheduling, error)
	FindAllSchedulingByMedicationID(ctx context.Context, medicationID *uint64) ([]*Scheduling, error)
	FindAllSchedulingByMedicationIDAndStatus(ctx context.Context, medicationID *uint64, status string) ([]*Scheduling, error)
	FindSchedulingByMedicationIDAndDateRange(ctx context.Context, medicationID *uint64, startDate, endDate time.Time) ([]*Scheduling, error)
	UpdateScheduling(ctx context.Context, scheduling *Scheduling) (*Scheduling, error)
	DeleteScheduling(ctx context.Context, ID *uint64) error
}
