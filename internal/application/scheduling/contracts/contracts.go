package scheduling

import (
	"context"

	scheduling "github.com/cuida-me/mvp-backend/internal/application/scheduling/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	domain "github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
)

type DoneScheduling interface {
	Execute(ctx context.Context, id *uint64) (*scheduling.DoneSchedulingResponse, *apiErr.Message)
}

type GetScheduling interface {
	Execute(ctx context.Context, id *uint64) (*domain.Scheduling, error)
}

type GetWeekScheduling interface {
	Execute(ctx context.Context, patientID *uint64) ([]*scheduling.DailyScheduling, *apiErr.Message)
}

type ScheduleWeekMedication interface {
	Execute(ctx context.Context) *scheduling.JobResponse
}

type GetReport interface {
	Execute(ctx context.Context, patientID *uint64) ([][]string, *apiErr.Message)
}

type SchedulingService interface {
	CreateAllSchedulingByMedication(ctx context.Context, medication medication.Medication) error
}
