package caregiver

import (
	"context"
	"fmt"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type linkPatientDeviceUseCase struct {
	repository               caregiver.Repository
	patientSessionRepository patient.SessionRepository
	patientRepository        patient.Repository
	log                      log.Provider
	apiErr                   apiErr.Provider
}

func NewLinkPatientDeviceUseCase(
	repository caregiver.Repository,
	patientSessionRepository patient.SessionRepository,
	patientRepository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *linkPatientDeviceUseCase {
	return &linkPatientDeviceUseCase{
		repository:               repository,
		patientSessionRepository: patientSessionRepository,
		patientRepository:        patientRepository,
		log:                      log,
		apiErr:                   apiErr,
	}
}

func (u linkPatientDeviceUseCase) Execute(ctx context.Context, qrToken *string, caregiverID *uint64) (*dto.LinkPatientDeviceResponse, string, *apiErr.Message) {
	u.log.Info(ctx, "linking patient device", log.Body{
		"qrToken": qrToken,
	})

	caregiver, err := u.repository.FindCaregiverByID(ctx, caregiverID)
	if err != nil {
		u.log.Error(ctx, "error to find caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, "", u.apiErr.InternalServerError(err)
	}

	patientSession, err := u.patientSessionRepository.FindPatientByQrToken(ctx, *qrToken)
	if err != nil {
		u.log.Error(ctx, "error to find patient session", log.Body{
			"error": err.Error(),
		})
		return nil, "", u.apiErr.InternalServerError(err)
	}

	patient, err := u.patientRepository.FindPatientByID(ctx, caregiver.PatientID)
	if err != nil {
		u.log.Error(ctx, "error to find patient", log.Body{
			"error": err.Error(),
		})
		return nil, "", u.apiErr.InternalServerError(err)
	}

	jwt, err := commons.NewJwt(fmt.Sprintf("%s_%d", "patient", patient.ID))

	patientSession.Patient = patient
	patientSession.PatientID = &patient.ID
	patientSession.Token = jwt

	updatedPatientSession, err := u.patientSessionRepository.UpdatePatientSession(ctx, patientSession)
	if err != nil {
		u.log.Error(ctx, "error to update patient session", log.Body{
			"error": err.Error(),
		})
		return nil, "", u.apiErr.InternalServerError(err)
	}

	return &dto.LinkPatientDeviceResponse{
		Token:   updatedPatientSession.Token,
		Success: true,
	}, updatedPatientSession.SocketID, nil
}
