package patient

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type createPatientUseCase struct {
	repository          patient.Repository
	caregiverRepository caregiver.Repository
	log                 log.Provider
	apiErr              apiErr.Provider
}

func NewCreatePatientUseCase(
	repository patient.Repository,
	caregiverRepository caregiver.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *createPatientUseCase {
	return &createPatientUseCase{
		repository:          repository,
		caregiverRepository: caregiverRepository,
		log:                 log,
		apiErr:              apiErr,
	}
}

func (u createPatientUseCase) Execute(ctx context.Context, request *dto.CreatePatientRequest, caregiverID *uint64) (*dto.CreatePatientResponse, *apiErr.Message) {
	u.log.Info(ctx, "creating patient", log.Body{
		"name":          request.Name,
		"date_of_birth": request.BirthDate,
		"sex":           request.Sex,
		"caregiver_id":  *caregiverID,
	})

	caregiver, err := u.caregiverRepository.FindCaregiverByID(ctx, caregiverID)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest(err.Error(), err)
	}
	if caregiver.PatientID != nil {
		u.log.Error(ctx, "caregiver already have a patient", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest("caregiver already have a patient", err)
	}

	sex := domain.Sex(request.Sex)

	patient := &patient.Patient{
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Sex:       &sex,
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

	caregiver.PatientID = &created.ID
	caregiver.Patient = created

	_, err = u.caregiverRepository.UpdateCaregiver(ctx, caregiver)
	if err != nil {
		u.log.Error(ctx, "error to update caregiver", log.Body{
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
		if p.Sex != nil && *p.Sex == domain.MALE {
			// TODO: Implements default image
		} else if p.Sex != nil && *p.Sex == domain.FEMALE {
			// TODO: Implements default image
		} else {
		}
	} else {
		p.Avatar = *avatar
	}
}
