package caregiver

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
)

type Create interface {
	Execute(ctx context.Context, request *dto.CreateCaregiverRequest) (*dto.CreateCaregiverResponse, *apiErr.Message)
}

type GetCaregiver interface {
	Execute(ctx context.Context, id *uint64) (*dto.GetCaregiverResponse, *apiErr.Message)
}

type Delete interface {
	Execute(ctx context.Context, id *uint64) *apiErr.Message
}

type Update interface {
	Execute(ctx context.Context, request *dto.UpdateCaregiverRequest, caregiverID *uint64) (*dto.UpdateCaregiverResponse, *apiErr.Message)
}

type LinkPatientDevice interface {
	Execute(ctx context.Context, qrToken *string, caregiverID *uint64) (*dto.LinkPatientDeviceResponse, string, *apiErr.Message)
}
