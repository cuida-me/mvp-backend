package medication

import "context"

type Repository interface {
	CreateMedication(ctx context.Context, medication *Medication) (*Medication, error)
	FindMedicationByID(ctx context.Context, ID *uint64) (*Medication, error)
	UpdateMedication(ctx context.Context, medication *Medication) (*Medication, error)
	DeleteMedication(ctx context.Context, ID *uint64) error
}

type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, schedule *Schedule) (*Schedule, error)
	FindScheduleByID(ctx context.Context, ID *uint64) (*Schedule, error)
	UpdateSchedule(ctx context.Context, schedule *Schedule) (*Schedule, error)
	DeleteSchedule(ctx context.Context, ID *uint64) error
}

type TypeRepository interface {
	FindAllTypes(ctx context.Context) ([]*Type, error)
	FindTypeByID(ctx context.Context, ID *uint64) (*Type, error)
}
