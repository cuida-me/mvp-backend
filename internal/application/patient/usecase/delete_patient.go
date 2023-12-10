package patient

import (
	"context"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type deletePatientUseCase struct {
	repository          patient.Repository
	caregiverRepository caregiver.Repository
	log                 log.Provider
	apiErr              apiErr.Provider
}

func NewDeletePatientUseCase(
	repository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *deletePatientUseCase {
	return &deletePatientUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u deletePatientUseCase) Execute(ctx context.Context, id *uint64) *apiErr.Message {
	u.log.Info(ctx, "deleting patient", log.Body{
		"id": id,
	})

	caregiver, err := u.caregiverRepository.FindCaregiverByPatientID(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	caregiver.PatientID = nil

	_, err = u.caregiverRepository.UpdateCaregiver(ctx, caregiver)
	if err != nil {
		u.log.Error(ctx, "error to update caregiver", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	err = u.repository.DeletePatient(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to delete patient", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	return nil
}
