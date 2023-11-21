package medication

import (
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type CreateMedicationRequest struct {
	Name      string                   `json:"name"`
	TypeID    uint64                   `json:"type_id"`
	Avatar    string                   `json:"avatar"`
	Schedules []*CreateScheduleRequest `json:"schedules"`
}

type CreateMedicationResponse struct {
	ID        uint64                    `json:"id"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	Patient   *patient.Patient          `json:"patient"`
	Avatar    string                    `json:"avatar"`
	Schedules []*CreateScheduleResponse `json:"schedules"`
	Status    string                    `json:"status"`
}

func (c *CreateMedicationResponse) ToDTO(d *medication.Medication) {
	c.ID = d.ID
	c.Name = d.Name
	c.Type = d.Type.Name
	c.Avatar = d.Avatar
	c.Status = d.Status

	var schedules []*CreateScheduleResponse
	for _, schedule := range d.Schedules {
		schedules = append(schedules, &CreateScheduleResponse{
			ID:          schedule.ID,
			DailyOfWeek: &schedule.DailyOfWeek,
			Time:        schedule.Time,
			LiteralDay:  schedule.LiteralDay,
		})
	}

	c.Schedules = schedules
}
