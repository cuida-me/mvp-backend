package scheduling

import (
	"context"
	scheduling "github.com/cuida-me/mvp-backend/internal/application/scheduling/dto"
)

//type DoneScheduling interface {
//	Execute(ctx context.Context, id *uint64) *apiErr.Message
//}
//
//type GetWeekScheduling interface {
//	Execute(ctx context.Context, request *dto.UpdateCaregiverRequest, caregiverID *uint64) (*dto.UpdateCaregiverResponse, *apiErr.Message)
//}
//
//type UpdateWeekNotifications interface {
//	Execute(ctx context.Context, request *dto.UpdateCaregiverRequest, caregiverID *uint64) (*dto.UpdateCaregiverResponse, *apiErr.Message)
//}

type ScheduleWeekMedication interface {
	Execute(ctx context.Context) *scheduling.JobResponse
}
