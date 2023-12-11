package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	schedulingService "github.com/cuida-me/mvp-backend/internal/application/scheduling/contracts"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type updateMedicationUseCase struct {
	repository           medication.Repository
	typeRepository       medication.TypeRepository
	schedulingRepository scheduling.Repository
	scheduleRepository   medication.ScheduleRepository
	timeRepository       medication.TimeRepository
	schedulingService    schedulingService.SchedulingService
	log                  log.Provider
	apiErr               apiErr.Provider
}

func NewUpdateMedicationUseCase(
	repository medication.Repository,
	typeRepository medication.TypeRepository,
	schedulingRepository scheduling.Repository,
	scheduleRepository medication.ScheduleRepository,
	timeRepository medication.TimeRepository,
	schedulingService schedulingService.SchedulingService,
	log log.Provider,
	apiErr apiErr.Provider,
) *updateMedicationUseCase {
	return &updateMedicationUseCase{
		repository:           repository,
		typeRepository:       typeRepository,
		schedulingRepository: schedulingRepository,
		scheduleRepository:   scheduleRepository,
		timeRepository:       timeRepository,
		schedulingService:    schedulingService,
		log:                  log,
		apiErr:               apiErr,
	}
}

func (u updateMedicationUseCase) Execute(ctx context.Context, request *dto.UpdateMedicationRequest, medicationID, patientID *uint64) (*dto.UpdateMedicationResponse, *apiErr.Message) {
	u.log.Info(ctx, "updating medication", log.Body{
		"request":      request,
		"medicationID": medicationID,
		"patientID":    patientID,
	})

	medication, err := u.repository.FindMedicationByID(ctx, medicationID)
	if err != nil {
		u.log.Error(ctx, "error to find medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.NotFounded(err)
	}

	if medication.PatientID != *patientID {
		return nil, u.apiErr.Unauthorized("unauthorized")
	}

	quantityUpdated, apiError := u.updateMedication(ctx, medication, request)
	if apiError != nil {
		return nil, apiError
	}

	scheduleUpdated, apiError := u.updateSchedules(ctx, medication, request)
	if apiError != nil {
		return nil, apiError
	}

	timeUpdated, apiError := u.updateTimes(ctx, medication, request)
	if apiError != nil {
		return nil, apiError
	}

	updated, err := u.repository.UpdateMedication(ctx, medication)
	if err != nil {
		u.log.Error(ctx, "error to update medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	if quantityUpdated || scheduleUpdated || timeUpdated {
		u.log.Info(ctx, "updating schedulings", log.Body{
			"medication_id": updated.ID,
		})

		apiError := u.updateScheduling(ctx, updated)
		if apiError != nil {
			return nil, apiError
		}

	}

	var response dto.UpdateMedicationResponse

	response.ToDTO(updated)

	return &response, nil
}

func (u updateMedicationUseCase) updateMedication(ctx context.Context, medication *medication.Medication, request *dto.UpdateMedicationRequest) (bool, *apiErr.Message) {
	anyQuantityUpdated := false

	if medication.Name != request.Name && request.Name != "" {
		medication.Name = request.Name
	}

	if medication.TypeID != request.TypeID && request.TypeID != 0 {
		newType, err := u.typeRepository.FindTypeByID(ctx, &request.TypeID)
		if err != nil {
			u.log.Error(ctx, "error to find type", log.Body{
				"error": err.Error(),
			})
			return false, u.apiErr.NotFounded(err)
		}
		medication.Type = *newType
		medication.TypeID = newType.ID

		anyQuantityUpdated = true
	}

	if medication.Avatar != request.Avatar && request.Avatar != "" {
		medication.Avatar = request.Avatar
	}

	if medication.Quantity != request.Quantity && request.Quantity != 0 {
		medication.Quantity = request.Quantity
		anyQuantityUpdated = true
	}

	return anyQuantityUpdated, nil
}

func (u updateMedicationUseCase) updateSchedules(ctx context.Context, medicationSaved *medication.Medication, request *dto.UpdateMedicationRequest) (bool, *apiErr.Message) {
	anyUpdate := false
	for _, schedule := range request.Schedules {
		if schedule.ID != 0 {
			for _, medicationSchedule := range medicationSaved.Schedules {
				if medicationSchedule.ID == schedule.ID {
					if schedule.Enabled != nil && *schedule.Enabled != medicationSchedule.Enabled {
						medicationSchedule.Enabled = *schedule.Enabled

						updated, err := u.scheduleRepository.UpdateSchedule(ctx, medicationSchedule)
						if err != nil {
							u.log.Error(ctx, "error to update schedule", log.Body{
								"error":       err.Error(),
								"schedule_id": schedule.ID,
							})
							return false, u.apiErr.InternalServerError(err)
						}

						medicationSchedule = updated
						anyUpdate = true
					}
				}
			}
		}
	}

	return anyUpdate, nil
}

func (u updateMedicationUseCase) updateTimes(ctx context.Context, medicationSaved *medication.Medication, request *dto.UpdateMedicationRequest) (bool, *apiErr.Message) {
	anyUpdate := false

	if request.Times == nil {
		return anyUpdate, nil
	}

	for index, medicationTime := range medicationSaved.Times {
		if !commons.ContainsStr(*request.Times, medicationTime.Time) {
			err := u.timeRepository.DeleteTime(ctx, &medicationTime.ID)
			if err != nil {
				u.log.Error(ctx, "error to delete time", log.Body{
					"error": err.Error(),
				})
				return false, u.apiErr.InternalServerError(err)
			}
			anyUpdate = true

			medicationSaved.Times[index] = nil
		}
	}

	return anyUpdate, nil
}

func (u updateMedicationUseCase) updateScheduling(ctx context.Context, medicationSaved *medication.Medication) *apiErr.Message {
	scheduling, err := u.schedulingRepository.FindAllSchedulingByMedicationIDAndStatus(ctx, &medicationSaved.ID, scheduling.TODO)
	if err != nil {
		u.log.Error(ctx, "error to find scheduling", log.Body{
			"error": err.Error(),
		})
		return u.apiErr.InternalServerError(err)
	}

	if scheduling == nil {
		u.log.Error(ctx, "scheduling not found", log.Body{})
	}

	for _, s := range scheduling {
		err := u.scheduleRepository.DeleteSchedule(ctx, &s.ID)
		if err != nil {
			u.log.Error(ctx, "error to delete scheduling", log.Body{
				"error": err.Error(),
				"id":    s.ID,
			})
			return u.apiErr.InternalServerError(err)
		}
	}

	go u.schedulingService.CreateAllSchedulingByMedication(ctx, *medicationSaved)

	return nil
}
