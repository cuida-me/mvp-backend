package scheduling

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
)

const (
	DONE = "DONE"
	TODO = "TODO"
)

type Scheduling struct {
	ID                uint64
	Medication        *medication.Medication
	MedicationID      uint64 `gorm:"foreignKey:MedicationID"`
	Dosage            string
	Quantity          int
	MedicationType    string
	MedicationTime    *time.Time
	MedicationTakenAt *time.Time
	Avatar            string
	Status            string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
