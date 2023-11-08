package patient

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type refreshSessionQRUseCase struct {
	repository patient.SessionRepository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewRefreshSessionQRUseCase(
	repository patient.SessionRepository,
	log log.Provider,
	apiErr apiErr.Provider,
) *refreshSessionQRUseCase {
	return &refreshSessionQRUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u refreshSessionQRUseCase) Execute(ctx context.Context, request *dto.RefreshSessionQRRequest, socketID string) (*dto.RefreshSessionQRResponse, *apiErr.Message) {
	session, err := u.repository.FindPatientByQrToken(ctx, request.OldQR)
	if err != nil {
		u.log.Error(ctx, "error getting patient session", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	token := commons.GenerateToken(40)

	session.QrToken = token

	if session.SocketID != socketID {
		session.SocketID = socketID
	}

	updated, err := u.repository.UpdatePatientSession(ctx, session)

	if err != nil {
		u.log.Error(ctx, "error updating patient session", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	return &dto.RefreshSessionQRResponse{
		Token: updated.QrToken,
	}, nil
}
