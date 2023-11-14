package caregiver

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type updateCaregiverUseCase struct {
	repository caregiver.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewUpdateCaregiverUseCase(
	repository caregiver.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *updateCaregiverUseCase {
	return &updateCaregiverUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u updateCaregiverUseCase) Execute(ctx context.Context, request *dto.UpdateCaregiverRequest, caregiverID *uint64) (*dto.UpdateCaregiverResponse, *apiErr.Message) {
	u.log.Info(ctx, "updating caregiver", log.Body{
		"request": request,
		"id":      caregiverID,
	})

	caregiver, err := u.repository.FindCaregiverByID(ctx, caregiverID)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	u.updateCaregiverDiff(request, caregiver)

	caregiver, err = u.repository.UpdateCaregiver(ctx, caregiver)
	if err != nil {
		u.log.Error(ctx, "error to update caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.UpdateCaregiverResponse
	response.ToDTO(caregiver)

	return &response, nil
}

func (u updateCaregiverUseCase) updateCaregiverDiff(request *dto.UpdateCaregiverRequest, caregiverSaved *caregiver.Caregiver) {
	if request.Name != nil {
		caregiverSaved.Name = *request.Name
	}

	if request.BirthDate != nil {
		caregiverSaved.BirthDate = request.BirthDate
	}

	if request.Avatar != nil {
		caregiverSaved.Avatar = *request.Avatar
	}

	if request.Sex != nil {
		caregiverSaved.Sex = domain.Sex(*request.Sex)
	}

	if request.Email != nil {
		caregiverSaved.Email = *request.Email
	}
}
