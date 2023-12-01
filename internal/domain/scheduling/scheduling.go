package scheduling

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
)

type Scheduling struct {
	ID                uint64
	Medication        *medication.Medication
	MedicationID      uint64 `gorm:"foreignKey:MedicationID"`
	Dosage            string
	Quantity          string
	MedicationTime    *time.Time
	MedicationTakenAt *time.Time
	Avatar            string
	Status            string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
