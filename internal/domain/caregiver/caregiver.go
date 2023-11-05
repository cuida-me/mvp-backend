package caregiver

import (
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"time"
)

const (
	CREATED = "CREATED"
)

type Caregiver struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	BirthDate *time.Time
	Avatar    string
	Sex       domain.Sex
	Email     string `gorm:"unique"`
	PatientID *uint64
	Patient   *patient.Patient `gorm:"foreignKey:PatientID"`
	Status    string           `gorm:"default:CREATED"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (c Caregiver) ToCaregiverDTO() *pb.Caregiver {
	return &pb.Caregiver{
		Id:     c.ID,
		Name:   c.Name,
		Avatar: c.Avatar,
		Birthdate: &pb.Date{
			Year:  int32(c.BirthDate.Year()),
			Month: int32(c.BirthDate.Month()),
			Day:   int32(c.BirthDate.Day()),
		},
		Sex:    pb.Sex(c.Sex),
		Email:  c.Email,
		Status: c.Status,
	}
}

func (c Caregiver) ToCaregiverFullDTO() *pb.CaregiverFull {
	caregiver := &pb.CaregiverFull{
		Id:     c.ID,
		Name:   c.Name,
		Avatar: c.Avatar,
		Birthdate: &pb.Date{
			Year:  int32(c.BirthDate.Year()),
			Month: int32(c.BirthDate.Month()),
			Day:   int32(c.BirthDate.Day()),
		},
		Sex:    pb.Sex(c.Sex),
		Email:  c.Email,
		Status: c.Status,
	}

	if c.Patient != nil {
		caregiver.Patient = c.Patient.ToPatientDTO()
	}

	return caregiver
}
