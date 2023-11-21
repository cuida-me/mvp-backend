package scheduling

import (
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/medication"
)

type Scheduling struct {
	ID                uint64
	Medication        *medication.Medication
	Dosage            string
	Quantity          string
	MedicationTime    *time.Time
	MedicationTakenAt *time.Time
	Status            string
	StatusDetail      string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
