package caregiver

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type getCaregiverUseCase struct {
	repository caregiver.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewGetCaregiverUseCase(
	repository caregiver.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *getCaregiverUseCase {
	return &getCaregiverUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u getCaregiverUseCase) Execute(ctx context.Context, id *uint64) (*dto.GetCaregiverResponse, *apiErr.Message) {
	u.log.Info(ctx, "getting caregiver", log.Body{
		"id": id,
	})

	caregiver, err := u.repository.FindCaregiverByID(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.GetCaregiverResponse
	response.ToDTO(caregiver)

	return &response, nil
}
