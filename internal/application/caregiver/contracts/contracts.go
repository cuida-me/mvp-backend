package caregiver

import (
	"context"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
)

type Create interface {
	Execute(ctx context.Context, request *dto.CreateCaregiverRequest) (*dto.CreateCaregiverResponse, *apiErr.Message)
}
