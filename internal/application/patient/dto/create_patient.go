package patient

import (
	"time"

	domain2 "github.com/cuida-me/mvp-backend/internal/domain"
	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type CreatePatientRequest struct {
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birth_date"`
	Avatar    *string    `json:"avatar"`
	Sex       int        `json:"sex"`
}

type CreatePatientResponse struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	BirthDate *time.Time   `json:"birth_date"`
	Avatar    string       `json:"avatar"`
	Sex       *domain2.Sex `json:"sex"`
	Status    string       `json:"status"`
}

func (p *CreatePatientResponse) ToDTO(d *domain.Patient) {
	p.ID = d.ID
	p.Name = d.Name
	p.BirthDate = d.BirthDate
	p.Avatar = d.Avatar
	p.Sex = d.Sex
	p.Status = d.Status
}
