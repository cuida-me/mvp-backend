package patient

import (
	"time"

	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type UpdatePatientRequest struct {
	Name      string
	BirthDate *time.Time
	Avatar    string
	Sex       *int
}

type UpdatePatientResponse struct {
	ID        uint64
	Name      string
	BirthDate *time.Time
	Avatar    string
	Sex       string
	Status    string
}

func (p *UpdatePatientResponse) ToDTO(d *domain.Patient) {
	p.ID = d.ID
	p.Name = d.Name
	p.BirthDate = d.BirthDate
	p.Avatar = d.Avatar
	p.Sex = d.Sex.String()
	p.Status = d.Status
}
