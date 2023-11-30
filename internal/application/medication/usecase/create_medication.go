package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type CreateMedicationUseCase struct {
	repository         medication.Repository
	scheduleRepository medication.ScheduleRepository
	typeRepository     medication.TypeRepository
	patientRepository  patient.Repository
	log                log.Provider
	apiErr             apiErr.Provider
}

func NewCreateMedicationUseCase(
	repository medication.Repository,
	scheduleRepository medication.ScheduleRepository,
	typeRepository medication.TypeRepository,
	patientRepository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *CreateMedicationUseCase {
	return &CreateMedicationUseCase{
		repository:         repository,
		scheduleRepository: scheduleRepository,
		typeRepository:     typeRepository,
		patientRepository:  patientRepository,
		log:                log,
		apiErr:             apiErr,
	}
}

func (u CreateMedicationUseCase) Execute(ctx context.Context, request *dto.CreateMedicationRequest, patientID *uint64) (*dto.CreateMedicationResponse, *apiErr.Message) {
	u.log.Info(ctx, "creating medication", log.Body{
		"request":   request,
		"patientID": patientID,
	})

	patient, err := u.patientRepository.FindPatientByID(ctx, patientID)
	if err != nil {
		u.log.Error(ctx, "error to find patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	medicationType, err := u.typeRepository.FindTypeByID(ctx, &request.TypeID)
	if err != nil {
		u.log.Error(ctx, "error to find type", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest(err.Error(), err)
	}

	schedules := make([]*medication.MedicationSchedule, 0)

	for _, schedule := range request.Schedules {
		if schedule.DailyOfWeek == nil {
			return nil, u.apiErr.BadRequest("daily of week is required", nil)
		}

		times := make([]*medication.MedicationScheduleTime, 0)
		for _, time := range schedule.Times {
			times = append(times, &medication.MedicationScheduleTime{
				Time: time,
			})
		}

		schedules = append(schedules, &medication.MedicationSchedule{
			DailyOfWeek: *schedule.DailyOfWeek,
			Times:       times,
			LiteralDay:  domain.DailyOfWeek(*schedule.DailyOfWeek).String(),
			Enabled:     true,
		})
	}

	medication := &medication.Medication{
		Name:      request.Name,
		Avatar:    request.Avatar,
		PatientID: patient.ID,
		Type:      *medicationType,
		Status:    medication.CREATED,
		Schedules: schedules,
		Dosage:    request.Dosage,
		Quantity:  request.Quantity,
	}

	created, err := u.repository.CreateMedication(ctx, medication)
	if err != nil {
		u.log.Error(ctx, "error to create medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.CreateMedicationResponse

	response.ToDTO(created)

	return &response, nil
}