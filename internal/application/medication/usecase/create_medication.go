package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	schedulingService "github.com/cuida-me/mvp-backend/internal/application/scheduling/contracts"
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
	schedulingService  schedulingService.SchedulingService
	log                log.Provider
	apiErr             apiErr.Provider
}

func NewCreateMedicationUseCase(
	repository medication.Repository,
	scheduleRepository medication.ScheduleRepository,
	typeRepository medication.TypeRepository,
	patientRepository patient.Repository,
	schedulingService schedulingService.SchedulingService,
	log log.Provider,
	apiErr apiErr.Provider,
) *CreateMedicationUseCase {
	return &CreateMedicationUseCase{
		repository:         repository,
		scheduleRepository: scheduleRepository,
		typeRepository:     typeRepository,
		patientRepository:  patientRepository,
		schedulingService:  schedulingService,
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

	times := make([]*medication.MedicationTime, 0)
	for _, time := range request.Times {
		times = append(times, &medication.MedicationTime{
			Time: time,
		})
	}

	schedules := make([]*medication.MedicationSchedule, 0)

	for _, schedule := range request.Schedules {
		if schedule.DailyOfWeek == nil {
			return nil, u.apiErr.BadRequest("daily of week is required", nil)
		}

		schedules = append(schedules, &medication.MedicationSchedule{
			DailyOfWeek: *schedule.DailyOfWeek,
			LiteralDay:  domain.DailyOfWeek(*schedule.DailyOfWeek).String(),
			Enabled:     schedule.Enabled,
		})
	}

	medication := &medication.Medication{
		Name:      request.Name,
		Avatar:    request.Avatar,
		PatientID: patient.ID,
		Type:      *medicationType,
		Status:    medication.CREATED,
		Times:     times,
		Schedules: schedules,
		Quantity:  request.Quantity,
	}

	created, err := u.repository.CreateMedication(ctx, medication)
	if err != nil {
		u.log.Error(ctx, "error to create medication", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	err = u.schedulingService.CreateAllSchedulingByMedication(ctx, *created)
	if err != nil {
		u.log.Error(ctx, "error to create scheduling", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.CreateMedicationResponse

	response.ToDTO(created)

	return &response, nil
}
