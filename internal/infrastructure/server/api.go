package server

import (
	caregiver "github.com/cuida-me/mvp-backend/internal/application/caregiver/usecase"
	medication "github.com/cuida-me/mvp-backend/internal/application/medication/usecase"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/usecase"
	scheduling "github.com/cuida-me/mvp-backend/internal/application/scheduling/job"
	schedulingService "github.com/cuida-me/mvp-backend/internal/application/scheduling/service"
	schedulingUseCase "github.com/cuida-me/mvp-backend/internal/application/scheduling/usecase"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/auth/firebase"
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
	firebase, err := firebase.GetConnection()
	if err != nil {
		return err
	}

	connection, err := mysql.GetConnection(BootstrapDatabase(*a.Cfg))
	if err != nil {
		return err
	}

	logger := jsonlogs.New(a.Cfg.LogLevel, logcontext.GetCtxValues)
	apiErrors := apierr.New()

	// Repositories
	patientRepository := repository.NewPatientRepository(connection)
	caregiverRepository := repository.NewCaregiverRepository(connection)
	patientSessionRepository := repository.NewPatientSessionRepository(connection)
	medicationRepository := repository.NewMedicationRepository(connection)
	medicationScheduleRepository := repository.NewMedicationScheduleRepository(connection)
	medicationTypeRepository := repository.NewMedicationTypeRepository(connection)
	medicationTimeRepository := repository.NewMedicationTimeRepository(connection)
	schedulingRepository := repository.NewSchedulingRepository(connection)

	// Services
	schedulingService := schedulingService.NewSchedulingService(schedulingRepository, logger, apiErrors)

	// UseCases
	createPatientUseCase := patient.NewCreatePatientUseCase(patientRepository, caregiverRepository, logger, apiErrors)
	getPatientUseCase := patient.NewGetPatientUseCase(patientRepository, logger, apiErrors)
	deletePatientUseCase := patient.NewDeletePatientUseCase(patientRepository, logger, apiErrors)
	updatePatientUseCase := patient.NewUpdatePatientUseCase(patientRepository, logger, apiErrors)

	newPatientSessionUseCase := patient.NewPatientSessionUseCase(patientSessionRepository, logger, apiErrors)
	refreshSessionQrUseCase := patient.NewRefreshSessionQRUseCase(patientSessionRepository, logger, apiErrors)

	createCaregiverUseCase := caregiver.NewCreateCaregiverUseCase(caregiverRepository, firebase, logger, apiErrors)
	getCaregiverUseCase := caregiver.NewGetCaregiverUseCase(caregiverRepository, logger, apiErrors)
	deleteCaregiverUseCase := caregiver.NewDeleteCaregiverUseCase(caregiverRepository, patientRepository, logger, apiErrors)
	updateCaregiverUseCase := caregiver.NewUpdateCaregiverUseCase(caregiverRepository, logger, apiErrors)
	linkPatientDeviceUseCase := caregiver.NewLinkPatientDeviceUseCase(caregiverRepository, patientSessionRepository, patientRepository, logger, apiErrors)

	createMedicationUseCase := medication.NewCreateMedicationUseCase(medicationRepository, medicationScheduleRepository, medicationTypeRepository, patientRepository, schedulingService, logger, apiErrors)
	getMedicationUseCase := medication.NewGetMedicationUseCase(medicationRepository, logger, apiErrors)
	deleteMedicationUseCase := medication.NewDeleteMedicationUseCase(medicationRepository, medicationScheduleRepository, medicationTimeRepository, schedulingRepository, logger, apiErrors)
	getMedicationTypes := medication.NewGetMedicationTypesUseCase(medicationTypeRepository, logger, apiErrors)
	updateMedicationUseCase := medication.NewUpdateMedicationUseCase(medicationRepository, medicationTypeRepository, schedulingRepository, medicationScheduleRepository, medicationTimeRepository, schedulingService, logger, apiErrors)

	doneSchedulingUseCase := schedulingUseCase.NewDoneSchedulingUseCase(schedulingRepository, logger, apiErrors)
	getSchedulingUseCase := schedulingUseCase.NewGetScheduling(schedulingRepository)
	getWeekSchedulingUsecase := schedulingUseCase.NewGetWeekSchedulingUseCase(schedulingRepository, medicationRepository, logger, apiErrors)

	job := scheduling.NewScheduleWeekMedicationJob(schedulingRepository, patientRepository, medicationRepository, schedulingService, logger, apiErrors)

	// Websocket
	websocket := socket_io.NewWebsocketConnection(newPatientSessionUseCase, refreshSessionQrUseCase)

	// Middlewares
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))
	a.Router.Use(middlewares.EnsureAuth(logger, caregiverRepository, firebase))
	a.Router.Use(middlewares.RequestLogger(logger))
	a.Router.Use(middlewares.HandleRequestID())

	session := websocket.SocketConnection()

	// Routes
	a.Router.HandleFunc("/ping", handler.Ping()).Methods("GET")

	a.Router.HandleFunc("/job/schedule-week-medication", handler.ScheduleWeekMedication(job)).Methods("POST")

	a.Router.HandleFunc("/patient", handler.CreatePatient(createPatientUseCase)).Methods("POST")
	a.Router.HandleFunc("/patient", handler.GetPatient(getPatientUseCase)).Methods("GET")
	a.Router.HandleFunc("/patient", handler.DeletePatient(deletePatientUseCase)).Methods("DELETE")
	a.Router.HandleFunc("/patient", handler.UpdatePatient(updatePatientUseCase)).Methods("PUT")

	a.Router.HandleFunc("/caregiver", handler.CreateCaregiver(createCaregiverUseCase)).Methods("POST")
	a.Router.HandleFunc("/caregiver", handler.GetCaregiver(getCaregiverUseCase)).Methods("GET")
	a.Router.HandleFunc("/caregiver", handler.DeleteCaregiver(deleteCaregiverUseCase)).Methods("DELETE")
	a.Router.HandleFunc("/caregiver", handler.UpdateCaregiver(updateCaregiverUseCase)).Methods("PUT")
	a.Router.HandleFunc("/caregiver", handler.UpdateCaregiver(updateCaregiverUseCase)).Methods("PUT")

	a.Router.HandleFunc("/medication/types", handler.GetMedicationTypes(getMedicationTypes)).Methods("GET")
	a.Router.HandleFunc("/medication", handler.CreateMedication(createMedicationUseCase)).Methods("POST")
	a.Router.HandleFunc("/medication/{medicationID}", handler.GetMedication(getMedicationUseCase)).Methods("GET")
	a.Router.HandleFunc("/medication/{medicationID}", handler.DeleteMedication(deleteMedicationUseCase)).Methods("DELETE")
	a.Router.HandleFunc("/medication/{medicationID}", handler.UpdateMedication(updateMedicationUseCase)).Methods("PUT")

	a.Router.HandleFunc("/scheduling/{schedulingID}", handler.DoneScheduling(doneSchedulingUseCase)).Methods("PUT")
	a.Router.HandleFunc("/scheduling/{schedulingID}", handler.GetScheduling(getSchedulingUseCase)).Methods("GET")
	a.Router.HandleFunc("/scheduling/week", handler.GetWeekScheduling(getWeekSchedulingUsecase)).Methods("GET")

	a.Router.HandleFunc("/caregiver/patient/device/{qr_token}", handler.LinkPatientDevice(linkPatientDeviceUseCase, session)).Methods("POST")

	a.Router.Handle("/socket.io/", session)

	return nil
}

func BootstrapDatabase(config Config) *mysql.ConnectionData {
	connection := &mysql.ConnectionData{}

	switch {
	case config.Environment.IsBeta():
		return connection.SetupBetaConnectionData(config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabaseSchema, config.DatabaseUrl)

	case config.Environment.IsProduction():
		return connection.SetupProdConnectionData()

	case config.Environment.IsLocal():
		return connection.SetupLocalConnectionData(config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabaseSchema, config.DatabaseUrl)

	case config.Environment.IsTest():
		return connection.SetupTestConnectionData()
	}

	return nil
}
