package scheduling

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
)

type DoneSchedulingResponse struct {
	ID                uint64     `json:"id"`
	Dosage            string     `json:"dosage"`
	Quantity          int        `json:"quantity"`
	Avatar            string     `json:"avatar"`
	MedicationTime    *time.Time `json:"medication_time"`
	MedicationTakenAt *time.Time `json:"medication_taken_at"`
	Status            string     `json:"status"`
}

func (d *DoneSchedulingResponse) ToDTO(s *scheduling.Scheduling) {
	d.ID = s.ID
	d.Dosage = s.Dosage
	d.Quantity = s.Quantity
	d.Avatar = s.Avatar
	d.MedicationTime = s.MedicationTime
	d.MedicationTakenAt = s.MedicationTakenAt
	d.Status = s.Status
}
