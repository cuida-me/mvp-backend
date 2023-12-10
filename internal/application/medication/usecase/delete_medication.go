package medication

import (
	"context"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type deleteMedicationUseCase struct {
	repository           medication.Repository
	scheduleRepository   medication.ScheduleRepository
	timeRepository       medication.TimeRepository
	schedulingRepository scheduling.Repository
	log                  log.Provider
	apiErr               apiErr.Provider
}

func NewDeleteMedicationUseCase(
	repository medication.Repository,
	scheduleRepository medication.ScheduleRepository,
	timeRepository medication.TimeRepository,
	schedulingRepository scheduling.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *deleteMedicationUseCase {
	return &deleteMedicationUseCase{
		repository:           repository,
		scheduleRepository:   scheduleRepository,
		schedulingRepository: schedulingRepository,
		timeRepository:       timeRepository,
		log:                  log,
		apiErr:               apiErr,
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

	for _, time := range medication.Times {
		err = u.timeRepository.DeleteTime(ctx, &time.ID)
		if err != nil {
			u.log.Error(ctx, "error to delete time of medication", log.Body{
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

	schedulings, err := u.schedulingRepository.FindAllSchedulingByMedicationIDAndStatus(ctx, medicationID, scheduling.TODO)
	if err != nil {
		u.log.Error(ctx, "error to find schedulings", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	for _, scheduling := range schedulings {
		err = u.schedulingRepository.DeleteScheduling(ctx, &scheduling.ID)
		if err != nil {
			u.log.Error(ctx, "error to delete scheduling", log.Body{
				"error": err.Error(),
			})
			return u.apiErr.InternalServerError(err)
		}
	}

	return nil
}
