package medication

import (
	"context"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type deleteMedicationUseCase struct {
	repository         medication.Repository
	scheduleRepository medication.ScheduleRepository
	log                log.Provider
	apiErr             apiErr.Provider
}

func NewDeleteMedicationUseCase(
	repository medication.Repository,
	scheduleRepository medication.ScheduleRepository,
	log log.Provider,
	apiErr apiErr.Provider,
) *deleteMedicationUseCase {
	return &deleteMedicationUseCase{
		repository:         repository,
		scheduleRepository: scheduleRepository,
		log:                log,
		apiErr:             apiErr,
	}
}

func (u deleteMedicationUseCase) Execute(ctx context.Context, medicationID, patientID *uint64) *apiErr.Message {
	u.log.Info(ctx, "deleting medication", log.Body{
		"medicationID": medicationID,
		"patientID":    patientID,
	})

	medication, err := u.repository.FindMedicationByID(ctx, medicationID)
	if err != nil {
		u.log.Error(ctx, "error to find medication", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.NotFounded(err)
	}
	if medication.PatientID != *patientID {
		return u.apiErr.Unauthorized("unauthorized")
	}

	for _, schedule := range medication.Schedules {
		err = u.scheduleRepository.DeleteSchedule(ctx, &schedule.ID)
		if err != nil {
			u.log.Error(ctx, "error to delete schedule", log.Body{
				"error": err.Error(),
			})
			return u.apiErr.InternalServerError(err)
		}
	}

	err = u.repository.DeleteMedication(ctx, medicationID)
	if err != nil {
		u.log.Error(ctx, "error to delete medication", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	return nil
}
