package caregiver

import (
	"context"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type deleteCaregiverUseCase struct {
	repository        caregiver.Repository
	patientRepository patient.Repository
	log               log.Provider
	apiErr            apiErr.Provider
}

func NewDeleteCaregiverUseCase(
	repository caregiver.Repository,
	patientRepository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *deleteCaregiverUseCase {
	return &deleteCaregiverUseCase{
		repository:        repository,
		patientRepository: patientRepository,
		log:               log,
		apiErr:            apiErr,
	}
}

func (u deleteCaregiverUseCase) Execute(ctx context.Context, id *uint64) *apiErr.Message {
	u.log.Info(ctx, "deleting caregiver", log.Body{
		"id": id,
	})

	caregiverSaved, err := u.repository.FindCaregiverByID(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	now := time.Now()

	caregiverSaved.Status = caregiver.CANCELLED
	caregiverSaved.UpdatedAt = &now

	if _, err := u.repository.UpdateCaregiver(ctx, caregiverSaved); err != nil {
		u.log.Error(ctx, "error to delete caregiver", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	if caregiverSaved.PatientID != nil {
		caregiverSaved.Patient.Status = caregiver.CANCELLED
		caregiverSaved.Patient.UpdatedAt = &now

		_, err := u.patientRepository.UpdatePatient(ctx, caregiverSaved.Patient)
		if err != nil {
			u.log.Error(ctx, "error to delete patient", log.Body{
				"error": err.Error(),
			})
			return u.apiErr.InternalServerError(err)
		}
	}

	return nil
}
