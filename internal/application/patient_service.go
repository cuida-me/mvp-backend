package application

import (
	"context"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"io"
)

const TOKEN_LENGTH = 50

type patientService struct {
	pb.UnimplementedPatientServiceServer
	repository patient.Repository
	log        log.Provider
}

func NewPatientService(repository patient.Repository, log log.Provider) *patientService {
	return &patientService{
		repository: repository,
		log:        log,
	}
}

func (s patientService) Create(ctx context.Context, request *pb.CreatePatientRequest) (*pb.Patient, error) {
	date, err := commons.ConvertToDate(
		request.Birthdate.Year,
		request.Birthdate.Month,
		request.Birthdate.Day)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to create patient")
	}

	s.log.Info(ctx, "creating patient", log.Body{
		"name":          request.Name,
		"date_of_birth": date,
		"sex":           request.Sex,
	})

	patient := &patient.Patient{
		Name:      request.Name,
		BirthDate: date,
		Sex:       domain.Sex(request.Sex),
		Status:    patient.CREATED,
	}

	s.resolvePatientAvatar(patient, request.Avatar)

	created, err := s.repository.CreatePatient(ctx, patient)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to create patient")
	}

	return created.ToPatientDTO(), err
}

func (s patientService) FindById(ctx context.Context, request *pb.FindPatientByIDRequest) (*pb.Patient, error) {
	s.log.Info(ctx, "get patient", log.Body{
		"id": request.Id,
	})

	patient, err := s.repository.FindPatientByID(ctx, &request.Id)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to find patient")
	}

	return patient.ToPatientDTO(), err
}

func (s patientService) Update(ctx context.Context, request *pb.UpdatePatientRequest) (*pb.Blank, error) {
	patientSaved, err := s.repository.FindPatientByID(ctx, &request.PatientId)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to get patient")
	}

	if patientSaved == nil {
		err := fmt.Errorf("patient %v not founded", request.PatientId)
		return &pb.Blank{}, s.handlerError(ctx, err, "patient not found")
	}

	date, err := commons.ConvertToDate(
		request.Birthdate.Year,
		request.Birthdate.Month,
		request.Birthdate.Day)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to update patient")
	}

	patient := &patient.Patient{
		Name:      request.Name,
		BirthDate: date,
		Sex:       domain.Sex(request.Sex),
		Avatar:    request.Avatar,
	}

	s.updatePatientDiff(patientSaved, patient)

	s.log.Info(ctx, "updating patient", log.Body{
		"id": request.PatientId,
	})

	_, err = s.repository.UpdatePatient(ctx, patientSaved)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to update patient")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

func (s patientService) Delete(ctx context.Context, request *pb.DeletePatientRequest) (*pb.Blank, error) {
	s.log.Info(ctx, "delete patient", log.Body{
		"id": request.Id,
	})

	err := s.repository.DeletePatient(ctx, &request.Id)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to delete patient")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

func (s patientService) NewSession(stream pb.PatientService_NewSessionServer) error {
	fmt.Println("aqui")
	for {
		fmt.Println("aqui 2")
		request, err := stream.Recv()
		fmt.Println("aqui 2")
		if err == io.EOF {
			fmt.Println("aqui 3")
			return nil
		}
		if err != nil {
			fmt.Println("aqui 4")
			return err
		}

		s.log.Info(stream.Context(), "new patient session", log.Body{
			"ip":        request.Ip,
			"device_id": request.DeviceId,
		})

		qrToken := commons.GenerateToken(TOKEN_LENGTH)

		if err := stream.Send(&pb.Session{
			Token:          qrToken,
			LoginCompleted: false,
		}); err != nil {
			return err
		}
	}
}

func (s patientService) resolvePatientAvatar(p *patient.Patient, avatar *string) {
	if avatar == nil {
		if p.Sex == domain.MALE {
			// TODO: Implements default image

		} else if p.Sex == domain.FEMALE {
			// TODO: Implements default image

		} else {

		}
	} else {
		p.Avatar = *avatar
	}
}

func (s patientService) updatePatientDiff(actual, new *patient.Patient) {
	if new.Name != "" {
		actual.Name = new.Name
	}

	if new.Sex != actual.Sex {
		actual.Sex = new.Sex
	}

	if new.BirthDate != nil {
		actual.BirthDate = new.BirthDate
	}

	if new.Avatar != "" {
		actual.Avatar = new.Avatar
	}
}

func (s patientService) handlerError(ctx context.Context, err error, title string) error {
	if err != nil {
		s.log.Error(ctx, title, log.Body{
			"error": err.Error(),
		})
		return fmt.Errorf("%s error: %s", title, err.Error())
	}

	return nil
}
