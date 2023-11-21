package medication

import (
	"context"

	dto "github.com/cuida-me/mvp-backend/internal/application/medication/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
)

type Create interface {
	Execute(ctx context.Context, request *dto.CreateMedicationRequest, patientID *uint64) (*dto.CreateMedicationResponse, *apiErr.Message)
}

type GetMedication interface {
	Execute(ctx context.Context, medicationID, patientID *uint64) (*dto.GetMedicationResponse, *apiErr.Message)
}

type GetMedicationTypes interface {
	Execute(ctx context.Context) ([]*medication.Type, *apiErr.Message)
}

type Delete interface {
	Execute(ctx context.Context, medicationID, patientID *uint64) *apiErr.Message
}

//
//type Update interface {
//	Execute(ctx context.Context, request *dto.UpdateCaregiverRequest, caregiverID *uint64) (*dto.UpdateCaregiverResponse, *apiErr.Message)
//}
