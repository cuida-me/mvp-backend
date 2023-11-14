package server

import (
	caregiver "github.com/cuida-me/mvp-backend/internal/application/caregiver/usecase"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/usecase"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/database/mysql"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/handler"
	middlewares "github.com/cuida-me/mvp-backend/internal/infrastructure/middleware"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/repository"
	socket_io "github.com/cuida-me/mvp-backend/internal/infrastructure/websocket/socket.io"
	logcontext "github.com/cuida-me/mvp-backend/pkg/context"
	apierr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log/jsonlogs"
	"github.com/gorilla/mux"
)

type Api struct {
	Cfg    *Config
	Router *mux.Router
}

func NewApi(cfg *Config) Api {
	return Api{
		Cfg:    cfg,
		Router: mux.NewRouter(),
	}
}

func (a *Api) Bootstrap() error {
	// Providers
	connection, err := mysql.GetConnection(BootstrapDatabase(a.Cfg.Environment))
	if err != nil {
		return err
	}
	logger := jsonlogs.New(a.Cfg.LogLevel, logcontext.GetCtxValues)
	apiErrors := apierr.New()

	// Repositories
	patientRepository := repository.NewPatientRepository(connection)
	caregiverRepository := repository.NewCaregiverRepository(connection)
	patientSessionRepository := repository.NewPatientSessionRepository(connection)

	// UseCases
	createPatientUseCase := patient.NewCreatePatientUseCase(patientRepository, caregiverRepository, logger, apiErrors)
	getPatientUseCase := patient.NewGetPatientUseCase(patientRepository, logger, apiErrors)
	deletePatientUseCase := patient.NewDeletePatientUseCase(patientRepository, logger, apiErrors)
	updatePatientUseCase := patient.NewUpdatePatientUseCase(patientRepository, logger, apiErrors)

	newPatientSessionUseCase := patient.NewPatientSessionUseCase(patientSessionRepository, logger, apiErrors)
	refreshSessionQrUseCase := patient.NewRefreshSessionQRUseCase(patientSessionRepository, logger, apiErrors)

	createCaregiverUseCase := caregiver.NewCreateCaregiverUseCase(caregiverRepository, logger, apiErrors)
	getCaregiverUseCase := caregiver.NewGetCaregiverUseCase(caregiverRepository, logger, apiErrors)
	deleteCaregiverUseCase := caregiver.NewDeleteCaregiverUseCase(caregiverRepository, patientRepository, logger, apiErrors)
	updateCaregiverUseCase := caregiver.NewUpdateCaregiverUseCase(caregiverRepository, logger, apiErrors)
	linkPatientDeviceUseCase := caregiver.NewLinkPatientDeviceUseCase(caregiverRepository, patientSessionRepository, patientRepository, logger, apiErrors)

	// Websocket
	websocket := socket_io.NewWebsocketConnection(newPatientSessionUseCase, refreshSessionQrUseCase)

	// Middlewares
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))
	a.Router.Use(middlewares.EnsureAuth(logger, patientRepository, caregiverRepository))

	session := websocket.SocketConnection()

	// Routes
	a.Router.HandleFunc("/ping", handler.Ping()).Methods("GET")

	a.Router.HandleFunc("/patient", handler.CreatePatient(createPatientUseCase)).Methods("POST")
	a.Router.HandleFunc("/patient", handler.GetPatient(getPatientUseCase)).Methods("GET")
	a.Router.HandleFunc("/patient", handler.DeletePatient(deletePatientUseCase)).Methods("DELETE")
	a.Router.HandleFunc("/patient", handler.UpdatePatient(updatePatientUseCase)).Methods("PUT")

	a.Router.HandleFunc("/caregiver", handler.CreateCaregiver(createCaregiverUseCase)).Methods("POST")
	a.Router.HandleFunc("/caregiver", handler.GetCaregiver(getCaregiverUseCase)).Methods("GET")
	a.Router.HandleFunc("/caregiver", handler.DeleteCaregiver(deleteCaregiverUseCase)).Methods("DELETE")
	a.Router.HandleFunc("/caregiver", handler.UpdateCaregiver(updateCaregiverUseCase)).Methods("PUT")
	a.Router.HandleFunc("/caregiver", handler.UpdateCaregiver(updateCaregiverUseCase)).Methods("PUT")

	a.Router.HandleFunc("/caregiver/patient-device/{qr_token}", handler.LinkPatientDevice(linkPatientDeviceUseCase, session)).Methods("POST")

	a.Router.Handle("/socket.io/", session)

	return nil
}

func BootstrapDatabase(environment Environment) *mysql.ConnectionData {
	connection := &mysql.ConnectionData{}

	switch {
	case environment.IsBeta():
		return connection.SetupBetaConnectionData()

	case environment.IsProduction():
		return connection.SetupProdConnectionData()

	case environment.IsLocal():
		return connection.SetupLocalConnectionData()

	case environment.IsTest():
		return connection.SetupTestConnectionData()
	}

	return nil
}
