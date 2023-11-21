package patient

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	apierr "github.com/cuida-me/mvp-backend/pkg/errors"
)

//go:generate mockgen -destination=./mocks.go -package=patient -source=./contracts.go

type Create interface {
	Execute(ctx context.Context, request *dto.CreatePatientRequest, caregiverID *uint64) (*dto.CreatePatientResponse, *apierr.Message)
}

type NewPatientSession interface {
	Execute(ctx context.Context, request *dto.NewPatientSessionRequest, socketID string) (*dto.NewPatientSessionResponse, *apierr.Message)
}

type RefreshSessionQR interface {
	Execute(ctx context.Context, request *dto.RefreshSessionQRRequest, socketID string) (*dto.RefreshSessionQRResponse, *apierr.Message)
}

type GetPatient interface {
	Execute(ctx context.Context, patientID *uint64) (*dto.GetPatientResponse, *apierr.Message)
}

type Update interface {
	Execute(ctx context.Context, request *dto.UpdatePatientRequest, patientID *uint64) (*dto.UpdatePatientResponse, *apierr.Message)
}

type Delete interface {
	Execute(ctx context.Context, patientID *uint64) *apierr.Message
}

//
//type Logout interface {
//	Execute(ctx context.Context, request *dto.LogoutPatientRequest) *apierr.Message
//}
//
//type GetHelp interface {
//	Execute(ctx context.Context, request *dto.GetHelpRequest) *apierr.Message
//}
