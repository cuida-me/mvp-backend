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
	Dosage    string
	Quantity  int
}

type UpdateScheduleRequest struct {
	ID          uint64   `json:"id"`
	DailyOfWeek *int     `json:"daily_of_week"`
	Times       []string `json:"times"`
	Enabled     bool     `json:"enabled"`
}

type UpdateScheduleResponse struct {
	ID          uint64   `json:"id"`
	DailyOfWeek *int     `json:"daily_of_week"`
	LiteralDay  string   `json:"literal_day"`
	Times       []string `json:"times"`
	Enabled     bool     `json:"enabled"`
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
		times := make([]string, 0)
		for _, time := range schedule.Times {
			times = append(times, time.Time)
		}

		schedules = append(schedules, &UpdateScheduleResponse{
			ID:          schedule.ID,
			DailyOfWeek: &schedule.DailyOfWeek,
			Times:       times,
			LiteralDay:  schedule.LiteralDay,
			Enabled:     schedule.Enabled,
		})
	}

	c.Schedules = schedules
}
