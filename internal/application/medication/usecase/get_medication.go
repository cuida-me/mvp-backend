package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type getMedicationUseCase struct {
	repository medication.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewGetMedicationUseCase(
	repository medication.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *getMedicationUseCase {
	return &getMedicationUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u getMedicationUseCase) Execute(ctx context.Context, medicationID, patientID *uint64) (*dto.GetMedicationResponse, *apiErr.Message) {
	u.log.Info(ctx, "getting medication", log.Body{
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

	var response dto.GetMedicationResponse
	response.ToDTO(medication)

	return &response, nil
}
