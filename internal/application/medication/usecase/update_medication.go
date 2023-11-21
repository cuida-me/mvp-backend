package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type updateMedicationUseCase struct {
	repository     medication.Repository
	typeRepository medication.TypeRepository
	log            log.Provider
	apiErr         apiErr.Provider
}

func NewUpdateMedicationUseCase(
	repository medication.Repository,
	typeRepository medication.TypeRepository,
	log log.Provider,
	apiErr apiErr.Provider,
) *updateMedicationUseCase {
	return &updateMedicationUseCase{
		repository:     repository,
		typeRepository: typeRepository,
		log:            log,
		apiErr:         apiErr,
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

	apiError := u.updateMedication(ctx, medication, request)
	if apiError != nil {
		return nil, apiError
	}

	u.updateSchedules(ctx, medication, request)

	updated, err := u.repository.UpdateMedication(ctx, medication)
	if err != nil {
		u.log.Error(ctx, "error to update medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.UpdateMedicationResponse

	response.ToDTO(updated)

	return &response, nil
}

func (u updateMedicationUseCase) updateMedication(ctx context.Context, medication *medication.Medication, request *dto.UpdateMedicationRequest) *apiErr.Message {
	if medication.Name != request.Name && request.Name != "" {
		medication.Name = request.Name
	}

	if medication.TypeID != request.TypeID && request.TypeID != 0 {
		newType, err := u.typeRepository.FindTypeByID(ctx, &request.TypeID)
		if err != nil {
			u.log.Error(ctx, "error to find type", log.Body{
				"error": err.Error(),
			})
			return u.apiErr.NotFounded(err)
		}
		medication.Type = *newType
		medication.TypeID = newType.ID
	}

	if medication.Avatar != request.Avatar && request.Avatar != "" {
		medication.Avatar = request.Avatar
	}

	// TODO: update quantity and dosage from schedulings
	if medication.Dosage != request.Dosage && request.Dosage != "" {
		medication.Dosage = request.Dosage
	}

	// TODO: update quantity and dosage from schedulings
	if medication.Quantity != request.Quantity && request.Quantity != 0 {
		medication.Quantity = request.Quantity
	}

	return nil
}

// TODO: Finalizar
func (u updateMedicationUseCase) updateSchedules(ctx context.Context, medicationSaved *medication.Medication, request *dto.UpdateMedicationRequest) {
	for _, schedule := range request.Schedules {
		if schedule.ID != 0 {
			for _, medicationSchedule := range medicationSaved.Schedules {
				if medicationSchedule.ID == schedule.ID {
					if schedule.Enabled != medicationSchedule.Enabled {
						medicationSchedule.Enabled = schedule.Enabled
					}
					times := make([]*medication.MedicationScheduleTime, 0)
					for _, time := range schedule.Times {
						medicationSchedule.Times = append(medicationSchedule.Times, &medication.MedicationScheduleTime{
							Time: time,
						})
					}
					medicationSchedule.Times = times

					// TODO: updade schedulings existant
				}
			}
		}
	}
}
