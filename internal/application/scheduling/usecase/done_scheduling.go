package usecase

import (
	"context"
	"time"

	dto "github.com/cuida-me/mvp-backend/internal/application/scheduling/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type doneSchedulingUseCase struct {
	repository scheduling.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewDoneSchedulingUseCase(
	repository scheduling.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *doneSchedulingUseCase {
	return &doneSchedulingUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u doneSchedulingUseCase) Execute(ctx context.Context, id *uint64) (*dto.DoneSchedulingResponse, *apiErr.Message) {
	u.log.Info(ctx, "done scheduling", log.Body{
		"id": id,
	})

	schedulingToDone, err := u.repository.FindSchedulingByID(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to done scheduling", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	now := time.Now().In(time.FixedZone("UTC-3", -3*3600))

	schedulingToDone.Status = scheduling.DONE
	schedulingToDone.MedicationTakenAt = &now
	schedulingToDone.UpdatedAt = &now

	updated, err := u.repository.UpdateScheduling(ctx, schedulingToDone)
	if err != nil {
		u.log.Error(ctx, "error to update done scheduling", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.DoneSchedulingResponse
	response.ToDTO(updated)

	return &response, nil
}
