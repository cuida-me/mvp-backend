package medication

import (
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type UpdateMedicationRequest struct {
	Name      string                   `json:"name"`
	TypeID    uint64                   `json:"type_id"`
	Avatar    string                   `json:"avatar"`
	Schedules []*UpdateScheduleRequest `json:"schedules"`
	Times     *[]string                `json:"times"`
	Dosage    string
	Quantity  int
}

type UpdateMedicationResponse struct {
	ID        uint64                    `json:"id"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	Patient   *patient.Patient          `json:"patient"`
	Avatar    string                    `json:"avatar"`
	Schedules []*UpdateScheduleResponse `json:"schedules"`
	Status    string                    `json:"status"`
	Times     []string                  `json:"times"`
	Dosage    string
	Quantity  int
}

type UpdateScheduleRequest struct {
	ID      uint64 `json:"id"`
	Enabled *bool  `json:"enabled"`
}

type UpdateScheduleResponse struct {
	ID          uint64 `json:"id"`
	DailyOfWeek *int   `json:"daily_of_week"`
	LiteralDay  string `json:"literal_day"`
	Enabled     bool   `json:"enabled"`
}

func (c *UpdateMedicationResponse) ToDTO(d *medication.Medication) {
	c.ID = d.ID
	c.Name = d.Name
	c.Type = d.Type.Name
	c.Avatar = d.Avatar
	c.Status = d.Status
	c.Dosage = d.Dosage
	c.Quantity = d.Quantity

	var schedules []*UpdateScheduleResponse
	for _, schedule := range d.Schedules {
		schedules = append(schedules, &UpdateScheduleResponse{
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
