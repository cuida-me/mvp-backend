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
			{Name: "COMPRIMIDO", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fcomprimido.jpg?alt=media&token=acf1c774-2858-4e2a-a541-0bb3f0fe73ca"},
			{Name: "CÁPSULA", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fcapsula.png?alt=media&token=e55b2a8e-932a-412f-9701-29de030e6200"},
			{Name: "ELIXIR", Dosage: "mL", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Felixir.png?alt=media&token=6f8c80c9-9809-43e9-88fe-3c98e24ffdbb"},
			{Name: "SOLUÇÃO", Dosage: "mL", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fsolucao.png?alt=media&token=63d9d486-d74d-4ef9-b625-3a9ec9e04cf7"},
			{Name: "POMADA", Dosage: "gramas", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fpomada.png?alt=media&token=aa78e4cf-0a16-4b45-8368-436ca5ba95e4"},
			{Name: "CREME", Dosage: "gramas", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fcreme.png?alt=media&token=745fd294-8eb5-47b6-94a5-b6a820351e26"},
			{Name: "INJETÁVEL", Dosage: "mL", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Finjetavel.jpg?alt=media&token=2f7099ab-ceeb-4d13-8f25-57d8c6c75f37"},
			{Name: "AEROSSOL", Dosage: "mL", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Faerosol.png?alt=media&token=7560d3e8-5143-4877-90cb-993110f16d47"},
			{Name: "ADESIVO TRANSDÉRMICO", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fadesivo.png?alt=media&token=b5fce990-868f-490e-9520-a1f7998170a3"},
			{Name: "SUPOSITÓRIO", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fsupo.png?alt=media&token=77fac706-f938-4042-92a7-33f690f9a74f"},
			{Name: "PÓ", Dosage: "gramas", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fpo.jpg?alt=media&token=ff87b7c4-c4a0-4b41-a28f-b2fecf7e8b06"},
			{Name: "EFERVESCENTE", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fefervescente.png?alt=media&token=49c6f884-5862-4933-b55f-894caab87724"},
			{Name: "GOTA", Dosage: "Gotas", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fgota.png?alt=media&token=0f5c19a4-3114-4b80-a9f6-3dac1b36d9a8"},
			{Name: "PASTILHA", Dosage: "Unidade", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fpastilha.png?alt=media&token=917632f9-fce7-4836-bd50-6cd2d43d6206"},
			{Name: "GEL", Dosage: "gramas", Avatar: "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fgel.png?alt=media&token=9617b609-f683-4670-9d6e-16ee448eb6f9"},
		}

		for _, tipo := range tiposMedicamentos {
			if err := client.Create(&tipo).Error; err != nil {
				return nil, err
			}
		}
	}

	return client, nil
}
