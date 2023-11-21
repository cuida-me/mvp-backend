package medication

import (
	"time"
)

const (
	CREATED = "CREATED"
)

type Medication struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	TypeID    uint64 `gorm:"foreignKey:TypeID"`
	Type      MedicationType
	PatientID uint64
	Avatar    string
	Dosage    string
	Quantity  int
	Schedules []*MedicationSchedule `gorm:"foreignKey:MedicationID"`
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type MedicationSchedule struct {
	ID           uint64 `gorm:"primaryKey"`
	MedicationID uint64 `gorm:"foreignKey:MedicationID"`
	Medication   *Medication
	DailyOfWeek  int                       `gorm:"column:daily_of_week; check:(daily_of_week >= 0) AND (daily_of_week <= 6)"`
	Times        []*MedicationScheduleTime `gorm:"foreignKey:MedicationScheduleID"`
	LiteralDay   string
	Enabled      bool `gorm:"default:true"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type MedicationScheduleTime struct {
	ID                   uint64 `gorm:"primaryKey"`
	MedicationScheduleID uint64 `gorm:"foreignKey:MedicationScheduleID"`
	MedicationSchedule   *MedicationSchedule
	Time                 string
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
}

type MedicationType struct {
	ID     uint64 `gorm:"primaryKey"`
	Name   string
	Avatar string
}
