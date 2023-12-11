package usecase

import (
	"context"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"time"
)

type getReportUseCase struct {
	repository           scheduling.Repository
	log                  log.Provider
	medicationRepository medication.Repository
	apiErr               apiErr.Provider
}

func NewGetReportUseCase(
	repository scheduling.Repository,
	log log.Provider,
	medicationRepository medication.Repository,
	apiErr apiErr.Provider,
) *getReportUseCase {
	return &getReportUseCase{
		repository:           repository,
		log:                  log,
		medicationRepository: medicationRepository,
		apiErr:               apiErr,
	}
}

func (u getReportUseCase) Execute(ctx context.Context, patientID *uint64) ([][]string, *apiErr.Message) {
	u.log.Info(ctx, "get report", log.Body{
		"patientID": patientID,
	})

	medications, err := u.medicationRepository.FindAllMedicationByPatientID(ctx, patientID)
	if err != nil {
		u.log.Error(ctx, "error to get medications", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var report [][]string
	report = append(report, u.getColumNames())

	for _, medication := range medications {
		schedulings, err := u.repository.FindAllSchedulingByMedicationIDAndStatus(ctx, &medication.ID, scheduling.DONE)
		if err != nil {
			u.log.Error(ctx, "error to get schedulings", log.Body{
				"error": err.Error(),
			})
			return nil, u.apiErr.InternalServerError(err)
		}

		for _, scheduling := range schedulings {
			var row []string
			row = append(row, medication.Name)
			row = append(row, medication.Type.Name)
			row = append(row, scheduling.Dosage)
			row = append(row, fmt.Sprintf("%d", scheduling.Quantity))
			row = append(row, scheduling.MedicationTime.In(time.FixedZone("UTC-3", -3*3600)).String())
			row = append(row, scheduling.MedicationTakenAt.In(time.FixedZone("UTC-3", -3*3600)).String())
			row = append(row, scheduling.Status)
			report = append(report, row)
		}
	}

	return report, nil
}

func (u getReportUseCase) getColumNames() []string {
	return []string{"Medicamento", "Tipo", "Dosagem", "Quantidade", "Horário", "Tomado em", "Status"}
}
