package patient

import (
	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
	"time"
)

type CreatePatientRequest struct {
	Name      string
	BirthDate *time.Time
	Avatar    *string
	Sex       int
}

type CreatePatientResponse struct {
	ID        uint64
	Name      string
	BirthDate *time.Time
	Avatar    string
	Sex       string
	Status    string
}

func (p *CreatePatientResponse) ToDTO(d *domain.Patient) {
	p.ID = d.ID
	p.Name = d.Name
	p.BirthDate = d.BirthDate
	p.Avatar = d.Avatar
	p.Sex = d.Sex.String()
	p.Status = d.Status
}
