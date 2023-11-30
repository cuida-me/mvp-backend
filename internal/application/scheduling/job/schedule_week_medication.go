package scheduling

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/scheduling/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type scheduleWeekMedicationJob struct {
	repository           scheduling.Repository
	patientRepository    patient.Repository
	medicationRepository medication.Repository
	log                  log.Provider
	apiErr               apiErr.Provider
}

func NewScheduleWeekMedicationJob(
	repository scheduling.Repository,
	patientRepository patient.Repository,
	medicationRepository medication.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *scheduleWeekMedicationJob {
	return &scheduleWeekMedicationJob{
		repository:           repository,
		patientRepository:    patientRepository,
		medicationRepository: medicationRepository,
		log:                  log,
		apiErr:               apiErr,
	}
}

func (u scheduleWeekMedicationJob) Execute(ctx context.Context) *dto.JobResponse {
	u.log.Info(nil, "scheduling week medication")

	var totalToProcess, processedWithSuccess, processedWithError int

	patients, err := u.patientRepository.FindAllPatientByStatus(ctx, patient.CREATED)
	if err != nil {
		u.log.Error(ctx, "error to find patients", log.Body{
			"error": err.Error(),
		})
		return &dto.JobResponse{
			TotalToProcess:       totalToProcess,
			ProcessedWithSuccess: processedWithSuccess,
			ProcessedWithError:   processedWithError,
			Error:                err,
		}
	}

	totalToProcess = len(patients)

	for _, patient := range patients {
		err := u.scheduleWeekMedication(ctx, patient)
		if err != nil {
			u.log.Error(ctx, "error to schedule week medication from patient", log.Body{
				"error":      err.Error(),
				"patient_id": patient.ID,
			})
			processedWithError++
			continue
		}
		processedWithSuccess++
	}

	return &dto.JobResponse{
		TotalToProcess:       totalToProcess,
		ProcessedWithSuccess: processedWithSuccess,
		ProcessedWithError:   processedWithError,
	}
}

func (u scheduleWeekMedicationJob) scheduleWeekMedication(ctx context.Context, patient *patient.Patient) error {
	medications, err := u.medicationRepository.FindAllMedicationByPatientID(ctx, &patient.ID)
	if err != nil {
		u.log.Error(ctx, "error to find medications from patient to schedule", log.Body{
			"error":      err.Error(),
			"patient_id": patient.ID,
		})
		return err
	}

	// TODO: Continuar aqui

	return nil
}
