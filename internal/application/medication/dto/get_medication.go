package medication

import (
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
)

type GetMedicationResponse struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Avatar    string `json:"avatar"`
	Dosage    string
	Quantity  int
	Times     []string                  `json:"times"`
	Schedules []*CreateScheduleResponse `json:"schedules"`
	Status    string                    `json:"status"`
}

func (c *GetMedicationResponse) ToDTO(d *medication.Medication) {
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
		})
	}

	times := make([]string, 0)
	for _, time := range d.Times {
		times = append(times, time.Time)
	}

	c.Schedules = schedules
	c.Times = times
}
