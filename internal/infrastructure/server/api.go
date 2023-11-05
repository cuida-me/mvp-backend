package server

import (
	"github.com/cuida-me/mvp-backend/internal/application"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/database"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/repository"
	logcontext "github.com/cuida-me/mvp-backend/pkg/context"
	"github.com/cuida-me/mvp-backend/pkg/log/jsonlogs"
	"google.golang.org/grpc"
)

type Api struct {
	Cfg    *Config
	Server *grpc.Server
}

func NewApi(cfg *Config) Api {
	return Api{
		Cfg: cfg,
	}
}

func (a *Api) Bootstrap() error {
	connection, err := database.GetConnection(BootstrapDatabase(a.Cfg.Environment))
	if err != nil {
		return err
	}

	logger := jsonlogs.New(a.Cfg.LogLevel, logcontext.GetCtxValues)

	// Providers

	// Repositories
	patientRepository := repository.NewPatientRepository(connection)
	caregiverRepository := repository.NewCaregiverRepository(connection)

	// Services
	patientService := application.NewPatientService(patientRepository, logger)
	caregiverService := application.NewCaregiverService(caregiverRepository, logger)

	// Middlewares

	//opts := []grpc.ServerOption{
	//	grpc.UnaryInterceptor(ensureValidToken),
	//	grpc.UnaryInterceptor(logInterceptor),
	//	grpc.UnaryInterceptor(logInterceptor),
	//	//grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	//}

	server := grpc.NewServer()

	pb.RegisterPatientServiceServer(server, patientService)
	pb.RegisterCaregiverServiceServer(server, caregiverService)

	a.Server = server

	return nil
}

func BootstrapDatabase(environment Environment) *database.ConnectionData {
	connection := &database.ConnectionData{}

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
