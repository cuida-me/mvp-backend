package medication

import (
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type CreateMedicationRequest struct {
	Name      string                   `json:"name"`
	TypeID    uint64                   `json:"type_id"`
	Avatar    string                   `json:"avatar"`
	Dosage    string                   `json:"dosage"`
	Quantity  int                      `json:"quantity"`
	Times     []string                 `json:"times"`
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
	Dosage    string                    `json:"dosage"`
	Quantity  int                       `json:"quantity"`
	Times     []string                  `json:"times"`
}

func (c *CreateMedicationResponse) ToDTO(d *medication.Medication) {
	c.ID = d.ID
	c.Name = d.Name
	c.Type = d.Type.Name
	c.Avatar = d.Avatar
	c.Status = d.Status
	c.Dosage = d.Dosage
	c.Quantity = d.Quantity

	var schedules []*CreateScheduleResponse
	for _, schedule := range d.Schedules {
		schedules = append(schedules, &CreateScheduleResponse{
			ID:          schedule.ID,
			DailyOfWeek: &schedule.DailyOfWeek,
			LiteralDay:  schedule.LiteralDay,
			Enabled:     schedule.Enabled,
		})
	}

	times := make([]string, 0)
	for _, time := range d.Times {
		times = append(times, time.Time)
	}

	c.Schedules = schedules
	c.Times = times
}
