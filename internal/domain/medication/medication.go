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
	Type      Type
	PatientID uint64
	Avatar    string
	Schedules []*Schedule `gorm:"foreignKey:MedicationID"`
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Schedule struct {
	ID           uint64 `gorm:"primaryKey"`
	MedicationID uint64 `gorm:"foreignKey:MedicationID"`
	Medication   *Medication
	DailyOfWeek  int `gorm:"column:daily_of_week; check:(daily_of_week >= 0) AND (daily_of_week <= 6)"`
	Time         string
	LiteralDay   string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type Type struct {
	ID     uint64 `gorm:"primaryKey"`
	Name   string
	Avatar string
}
