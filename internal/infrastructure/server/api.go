package server

import (
	"github.com/cuida-me/mvp-backend/internal/application"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/usecase"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/cache/redis"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/database/mysql"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/handler"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/repository"
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
	redisClient := redis.New()
	connection, err := mysql.GetConnection(BootstrapDatabase(a.Cfg.Environment))
	if err != nil {
		return err
	}

	logger := jsonlogs.New(a.Cfg.LogLevel, logcontext.GetCtxValues)

	apiErrors := apierr.New()

	// Providers

	// Repositories
	patientRepository := repository.NewPatientRepository(connection)
	caregiverRepository := repository.NewCaregiverRepository(connection)

	// Services
	_ = application.NewPatientService(patientRepository, logger)
	_ = application.NewCaregiverService(caregiverRepository, patientRepository, logger)

	// UseCases
	createPatientUseCase := patient.NewCreatePatientUseCase(patientRepository, logger, apiErrors)
	newPatientSessionUseCase := patient.NewPatientSessionUseCase(patientRepository, logger, apiErrors, redisClient)

	// Middlewares
	a.Router.Use(mux.CORSMethodMiddleware(a.Router))

	// Routes
	a.Router.HandleFunc("/ping", handler.Ping()).Methods("GET")
	a.Router.HandleFunc("/patient", handler.CreatePatient(createPatientUseCase)).Methods("POST")
	a.Router.HandleFunc("/patient/session", handler.NewPatientSession(newPatientSessionUseCase))

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
