package mysql

import (
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	_ "github.com/go-sql-driver/mysql" // need to load mysql driver on api
	"gorm.io/gorm"
)

const (
	minutesConnMaxLifetime = 2
	maxIdleConnections     = 50
	maxOpenConnections     = 100
)

func GetConnection(data *ConnectionData) (*gorm.DB, error) {
	if data == nil {
		return nil, fmt.Errorf("connection data is nil")
	}

	client, err := gorm.Open(data.toDialect(), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := client.DB()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * minutesConnMaxLifetime)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)

	client.AutoMigrate(&patient.Patient{})
	client.AutoMigrate(&caregiver.Caregiver{})
	client.AutoMigrate(&patient.PatientSession{})
	client.AutoMigrate(&medication.Medication{}, &medication.MedicationSchedule{}, &medication.MedicationScheduleTime{}, &medication.MedicationType{})
	client.AutoMigrate(scheduling.Scheduling{})

	return client, nil
}
