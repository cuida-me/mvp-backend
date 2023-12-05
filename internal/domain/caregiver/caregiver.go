package caregiver

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
)

const (
	CREATED   = "CREATED"
	CANCELLED = "CANCELLED"
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
	Uid       string           `gorm:"unique"`
	SocketID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
