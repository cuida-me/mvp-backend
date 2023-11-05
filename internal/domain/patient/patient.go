package patient

import (
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"time"
)

const (
	CREATED   = "CREATED"
	CANCELLED = "CANCELLED"
)

type Patient struct {
	ID             uint64
	Name           string
	BirthDate      *time.Time
	Avatar         string
	Sex            domain.Sex
	Status         string
	SessionID      *uint64
	PatientSession *Session `gorm:"foreignKey:SessionID"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type Session struct {
	ID        uint64
	PatientID uint64
	Patient   *Patient `gorm:"foreignKey:PatientID"`
	Token     string
	Status    string
	IP        string
	DeviceID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p Patient) ToPatientDTO() *pb.Patient {
	return &pb.Patient{
		Id:     p.ID,
		Name:   p.Name,
		Avatar: p.Avatar,
		Sex:    pb.Sex(p.Sex),
		Status: p.Status,
		Birthdate: &pb.Date{
			Year:  int32(p.BirthDate.Year()),
			Month: int32(p.BirthDate.Month()),
			Day:   int32(p.BirthDate.Day()),
		},
	}
}
