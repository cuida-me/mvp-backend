package patient

import (
	"time"
)

type PatientSession struct {
	ID        uint64
	PatientID *uint64
	Patient   *Patient `gorm:"foreignKey:PatientID",`
	Token     string
	Status    string
	IP        string
	DeviceID  string
	QrToken   string
	SocketID  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
