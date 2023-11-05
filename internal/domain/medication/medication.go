package medication

import (
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"time"
)

type Medication struct {
	ID           uint64
	Name         string
	Type         string
	Patient      *patient.Patient
	Avatar       string
	Schedules    []*Schedule
	Status       string
	StatusDetail string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type Schedule struct {
	ID          uint64
	Medication  *Medication
	DailyOfWeek string
	Time        string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
