package patient

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type updatePatientUseCase struct {
	repository patient.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewUpdatePatientUseCase(
	repository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *updatePatientUseCase {
	return &updatePatientUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u updatePatientUseCase) Execute(ctx context.Context, request *dto.UpdatePatientRequest, patientID *uint64) (*dto.UpdatePatientResponse, *apiErr.Message) {
	patientSaved, err := u.repository.FindPatientByID(ctx, patientID)
	if err != nil {
		u.log.Error(ctx, "error to find patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest(err.Error(), err)
	}

	u.updatePatientDiff(request, patientSaved)

	u.log.Info(ctx, "updating patient", log.Body{
		"id": patientID,
	})

	_, err = u.repository.UpdatePatient(ctx, patientSaved)
	if err != nil {
		u.log.Error(ctx, "error to update patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	var response dto.UpdatePatientResponse

	response.ToDTO(patientSaved)

	return &response, nil
}

func (u updatePatientUseCase) updatePatientDiff(request *dto.UpdatePatientRequest, patientSaved *patient.Patient) {
	if request.Name != "" {
		patientSaved.Name = request.Name
	}

	if request.BirthDate != nil {
		patientSaved.BirthDate = request.BirthDate
	}

	if request.Avatar != "" {
		patientSaved.Avatar = request.Avatar
	}

	if request.Sex != nil {
		patientSaved.Sex = domain.Sex(*request.Sex)
	}
}
