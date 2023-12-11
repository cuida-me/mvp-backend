package medication

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/patient"
)

const (
	CREATED = "CREATED"
)

type Medication struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	TypeID    uint64 `gorm:"foreignKey:TypeID"`
	Type      MedicationType
	PatientID uint64 `gorm:"foreignKey:PatientID"`
	Patient   *patient.Patient
	Avatar    string
	Quantity  int
	Schedules []*MedicationSchedule `gorm:"foreignKey:MedicationID"`
	Status    string
	Times     []*MedicationTime `gorm:"foreignKey:MedicationID"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type MedicationSchedule struct {
	ID           uint64 `gorm:"primaryKey"`
	MedicationID uint64 `gorm:"foreignKey:MedicationID"`
	Medication   *Medication
	DailyOfWeek  int `gorm:"column:daily_of_week; check:(daily_of_week >= 0) AND (daily_of_week <= 6)"`
	LiteralDay   string
	Enabled      bool
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type MedicationTime struct {
	ID           uint64 `gorm:"primaryKey"`
	MedicationID uint64 `gorm:"foreignKey:MedicationID"`
	Medication   *Medication
	Time         string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type MedicationType struct {
	ID     uint64 `gorm:"primaryKey"`
	Name   string
	Avatar string
	Dosage string
}
