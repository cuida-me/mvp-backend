package caregiver

import (
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"time"
)

const (
	MALE    = "male"
	FEMALE  = "female"
	CREATED = "CREATED"
)

type Caregiver struct {
	ID        uint64
	Name      string
	BirthDate *time.Time
	Avatar    string
	Sex       string
	Email     string
	Patient   *patient.Patient
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (c Caregiver) ToCaregiverDTO() *pb.Caregiver {
	return &pb.Caregiver{
		ID:     c.ID,
		Name:   c.Name,
		Avatar: c.Avatar,
		Birthdate: &pb.Date{
			Year:  int32(c.BirthDate.Year()),
			Month: int32(c.BirthDate.Month()),
			Day:   int32(c.BirthDate.Day()),
		},
		Sex:    c.Sex,
		Email:  c.Email,
		Status: c.Status,
	}
}

func (c Caregiver) ToCaregiverFullDTO() *pb.CaregiverFull {
	caregiver := &pb.CaregiverFull{
		ID:     c.ID,
		Name:   c.Name,
		Avatar: c.Avatar,
		Birthdate: &pb.Date{
			Year:  int32(c.BirthDate.Year()),
			Month: int32(c.BirthDate.Month()),
			Day:   int32(c.BirthDate.Day()),
		},
		Sex:    c.Sex,
		Email:  c.Email,
		Status: c.Status,
	}

	if c.Patient != nil {
		caregiver.Patient = c.Patient.ToPatientDTO()
	}

	return caregiver
}
