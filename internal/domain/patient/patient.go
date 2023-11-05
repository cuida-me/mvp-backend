package patient

import (
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"time"
)

const (
	MALE    = "male"
	FEMALE  = "female"
	CREATED = "CREATED"
)

type Patient struct {
	ID             uint64
	Name           string
	BirthDate      *time.Time
	Avatar         string
	Sex            string
	Status         string
	PatientSession *Session
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

type Session struct {
	ID        uint64
	Patient   *Patient
	Token     string
	Status    string
	IP        string
	DeviceID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p Patient) ToPatientDTO() *pb.Patient {
	return &pb.Patient{
		ID:     p.ID,
		Name:   p.Name,
		Avatar: p.Avatar,
		Sex:    p.Sex,
		Status: p.Status,
		Birthdate: &pb.Date{
			Year:  int32(p.BirthDate.Year()),
			Month: int32(p.BirthDate.Month()),
			Day:   int32(p.BirthDate.Day()),
		},
	}
}
