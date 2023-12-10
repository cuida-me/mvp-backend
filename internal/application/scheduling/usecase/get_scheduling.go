package usecase

import (
	"context"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
)

type getScheduling struct {
	repository scheduling.Repository
}

func NewGetScheduling(repository scheduling.Repository) *getScheduling {
	return &getScheduling{
		repository: repository,
	}
}

func (g getScheduling) Execute(ctx context.Context, id *uint64) (*scheduling.Scheduling, error) {
	scheduling, err := g.repository.FindSchedulingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	formatted := scheduling.MedicationTime.In(time.FixedZone("UTC-3", -3*3600))

	scheduling.MedicationTime = &formatted
	return scheduling, nil
}
