package patient

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type createPatientUseCase struct {
	repository patient.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewCreatePatientUseCase(
	repository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *createPatientUseCase {
	return &createPatientUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u createPatientUseCase) Execute(ctx context.Context, request *dto.CreatePatientRequest) (*dto.CreatePatientResponse, *apiErr.Message) {
	u.log.Info(ctx, "creating patient", log.Body{
		"name":          request.Name,
		"date_of_birth": request.BirthDate,
		"sex":           request.Sex,
	})

	patient := &patient.Patient{
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Sex:       domain.Sex(request.Sex),
		Status:    patient.CREATED,
	}

	u.resolvePatientAvatar(patient, request.Avatar)

	created, err := u.repository.CreatePatient(ctx, patient)
	if err != nil {
		u.log.Error(ctx, "error to create patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	response := dto.CreatePatientResponse{}

	response.ToDTO(created)

	return &response, nil
}

func (u createPatientUseCase) resolvePatientAvatar(p *patient.Patient, avatar *string) {
	if avatar == nil {
		if p.Sex == domain.MALE {
			// TODO: Implements default image

		} else if p.Sex == domain.FEMALE {
			// TODO: Implements default image

		} else {

		}
	} else {
		p.Avatar = *avatar
	}
}
