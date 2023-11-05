package application

import (
	"context"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

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
		Sex:       request.Sex,
		Avatar:    request.Avatar,
		Status:    patient.CREATED,
	}

	s.resolvePatientAvatar(patient)

	created, err := s.repository.CreatePatient(ctx, patient)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to create patient")
	}

	return created.ToPatientDTO(), err
}

func (s patientService) FindById(ctx context.Context, request *pb.FindPatientByIDRequest) (*pb.Patient, error) {
	s.log.Info(ctx, "get patient", log.Body{
		"id": request.ID,
	})

	patient, err := s.repository.FindPatientByID(ctx, &request.ID)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to find patient")
	}

	return patient.ToPatientDTO(), err
}

func (s patientService) Update(ctx context.Context, request *pb.UpdatePatientRequest) (*pb.Blank, error) {
	patientSaved, err := s.repository.FindPatientByID(ctx, &request.PatientID)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to get patient")
	}

	if patientSaved == nil {
		err := fmt.Errorf("patient %v not founded", request.PatientID)
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
		Sex:       request.Sex,
		Avatar:    request.Avatar,
	}

	s.updatePatientDiff(patientSaved, patient)

	s.log.Info(ctx, "updating patient", log.Body{
		"id": request.PatientID,
	})

	_, err = s.repository.UpdatePatient(ctx, patientSaved)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to update patient")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

func (s patientService) Delete(ctx context.Context, request *pb.DeletePatientRequest) (*pb.Blank, error) {
	s.log.Info(ctx, "delete patient", log.Body{
		"id": request.ID,
	})

	err := s.repository.DeletePatient(ctx, &request.ID)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to delete patient")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

//func (s patientService) NewSession(ctx context.Context, request *pb.NewPatientSessionRequest) (*pb.Session, error) {
//
//}

func (s patientService) resolvePatientAvatar(p *patient.Patient) {
	if p.Avatar == "" {
		if p.Sex == patient.MALE {
			// TODO: Implements default image

		} else if p.Sex == patient.FEMALE {
			// TODO: Implements default image

		}
	}
}

func (s patientService) updatePatientDiff(actual, new *patient.Patient) {
	if new.Name != "" {
		actual.Name = new.Name
	}

	if new.Sex != "" {
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
