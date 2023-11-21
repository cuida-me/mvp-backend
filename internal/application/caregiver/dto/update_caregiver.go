package caregiver

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
)

type UpdateCaregiverRequest struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	BirthDate *time.Time `json:"birth_date"`
	Sex       *int
	Avatar    *string `json:"avatar"`
}

type UpdateCaregiverResponse struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birth_date"`
	Avatar    string     `json:"avatar"`
	Sex       string
	Email     string `json:"email"`
	Status    string `json:"status"`
}

func (c *UpdateCaregiverResponse) ToDTO(d *caregiver.Caregiver) {
	c.ID = d.ID
	c.Name = d.Name
	c.BirthDate = d.BirthDate
	c.Avatar = d.Avatar
	c.Sex = d.Sex.String()
	c.Email = d.Email
	c.Status = d.Status
}
