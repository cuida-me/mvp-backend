package patient

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type newPatientSessionUseCase struct {
	repository patient.SessionRepository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewPatientSessionUseCase(
	repository patient.SessionRepository,
	log log.Provider,
	apiErr apiErr.Provider,
) *newPatientSessionUseCase {
	return &newPatientSessionUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u newPatientSessionUseCase) Execute(ctx context.Context, request *dto.NewPatientSessionRequest, socketID string) (*dto.NewPatientSessionResponse, *apiErr.Message) {
	u.log.Info(ctx, "new patient session", log.Body{
		"ip":        request.Ip,
		"device_id": request.DeviceID,
	})

	token := commons.GenerateToken(40)

	session, err := u.repository.CreatePatientSession(ctx, &patient.PatientSession{
		DeviceID: request.DeviceID,
		QrToken:  token,
		IP:       request.Ip,
		SocketID: socketID,
	})

	if err != nil {
		u.log.Error(ctx, "error creating patient session", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	return &dto.NewPatientSessionResponse{
		Token: session.QrToken,
	}, nil
}
