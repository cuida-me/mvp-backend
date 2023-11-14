package caregiver

import (
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"time"
)

type GetCaregiverResponse struct {
	ID        uint64                      `json:"id"`
	Name      string                      `json:"name"`
	BirthDate *time.Time                  `json:"birth_date"`
	Avatar    string                      `json:"avatar"`
	Sex       string                      `json:"sex"`
	Email     string                      `json:"email"`
	Status    string                      `json:"status"`
	Patient   *patient.GetPatientResponse `json:"patient"`
}

func (c *GetCaregiverResponse) ToDTO(d *caregiver.Caregiver) {
	c.ID = d.ID
	c.Name = d.Name
	c.BirthDate = d.BirthDate
	c.Avatar = d.Avatar
	c.Sex = d.Sex.String()
	c.Email = d.Email
	c.Status = d.Status
	if d.Patient != nil {
		p := &patient.GetPatientResponse{}
		p.ToDTO(d.Patient)
		c.Patient = p
	}
}
