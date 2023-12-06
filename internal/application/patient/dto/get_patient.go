package patient

import (
	domain2 "github.com/cuida-me/mvp-backend/internal/domain"
	"time"

	domain "github.com/cuida-me/mvp-backend/internal/domain/patient"
)

type GetPatientResponse struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	BirthDate *time.Time   `json:"birth_date"`
	Avatar    string       `json:"avatar"`
	Sex       *domain2.Sex `json:"sex"`
	Status    string       `json:"status"`
}

func (p *GetPatientResponse) ToDTO(d *domain.Patient) {
	p.ID = d.ID
	p.Name = d.Name
	p.BirthDate = d.BirthDate
	p.Avatar = d.Avatar
	p.Sex = d.Sex
	p.Status = d.Status
}
