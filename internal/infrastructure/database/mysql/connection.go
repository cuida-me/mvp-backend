package mysql

import (
	"fmt"
	"time"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/medication"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/domain/scheduling"
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
	client.AutoMigrate(&medication.Medication{}, &medication.MedicationSchedule{}, &medication.MedicationTime{}, &medication.MedicationType{})
	client.AutoMigrate(scheduling.Scheduling{})

	var count int64
	client.Model(&medication.MedicationType{}).Count(&count)

	if count == 0 {
		tiposMedicamentos := []medication.MedicationType{
			{Name: "COMPRIMIDO", Avatar: ""},
			{Name: "CÁPSULA", Avatar: ""},
			{Name: "DRÁGEA", Avatar: ""},
			{Name: "ELIXIR", Avatar: ""},
			{Name: "SUSPENSÃO", Avatar: ""},
			{Name: "SOLUÇÃO", Avatar: ""},
			{Name: "POMADA", Avatar: ""},
			{Name: "CREME", Avatar: ""},
			{Name: "INJETÁVEL", Avatar: ""},
			{Name: "AEROSSOL", Avatar: ""},
			{Name: "ADESIVO TRANSDÉRMICO", Avatar: ""},
			{Name: "SUPOSITÓRIO", Avatar: ""},
			{Name: "PÓ", Avatar: ""},
			{Name: "EFERVESCENTE", Avatar: ""},
			{Name: "GOTA", Avatar: ""},
			{Name: "PASTILHA", Avatar: ""},
		}

		for _, tipo := range tiposMedicamentos {
			if err := client.Create(&tipo).Error; err != nil {
				return nil, err
			}
		}
	}

	return client, nil
}
