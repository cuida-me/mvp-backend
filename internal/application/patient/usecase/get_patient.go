package patient

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type getPatientUseCase struct {
	repository patient.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewGetPatientUseCase(
	repository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *getPatientUseCase {
	return &getPatientUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u getPatientUseCase) Execute(ctx context.Context, id *uint64) (*dto.GetPatientResponse, *apiErr.Message) {
	u.log.Info(ctx, "get patient", log.Body{
		"id": id,
	})

	patient, err := u.repository.FindPatientByID(ctx, id)
	if err != nil {
		u.log.Error(ctx, "error to get patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	response := dto.GetPatientResponse{}

	response.ToDTO(patient)

	return &response, nil
}
