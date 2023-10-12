package application

import (
	"context"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/storage"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"os"
)

type patientService struct {
	pb.UnimplementedCreatePatientServiceServer
	repository patient.Repository
	log        log.Provider
	storage    storage.Provider
}

func NewPatientService(repository patient.Repository, log log.Provider, storage storage.Provider) *patientService {
	return &patientService{
		repository: repository,
		log:        log,
		storage:    storage,
	}
}

func (s patientService) CreatePatient(ctx context.Context, request *pb.CreatePatientRequest) (*pb.Patient, error) {
	time := request.DateOfBirth.AsTime()

	s.log.Info(ctx, "creating patient", log.Body{
		"name":          request.Name,
		"date_of_birth": time,
		"sex":           request.Sex,
	})

	patient := patient.Patient{
		Name:      request.Name,
		DateBirth: &time,
		Sex:       request.Sex,
	}

	created, err := s.repository.CreatePatient(ctx, &patient)
	if err != nil {
		s.log.Error(ctx, "error to create patient", log.Body{
			"error": err.Error(),
		})
		return nil, err
	}

	return &pb.Patient{
		ID:             created.ID,
		Name:           created.Name,
		DateOfBirth:    timestamppb.New(*created.DateBirth),
		Sex:            created.Sex,
		ProfilePicture: patient.ProfilePicture,
	}, err
}

func (s patientService) GetPatient(ctx context.Context, request *pb.GetPatientRequest) (*pb.Patient, error) {
	s.log.Info(ctx, "get patient", log.Body{
		"id": request.ID,
	})

	patient, err := s.repository.FindPatientByID(ctx, &request.ID)
	if err != nil {
		s.log.Error(ctx, "error to get patient", log.Body{
			"error": err.Error(),
		})
		return nil, err
	}

	return &pb.Patient{
		ID:             patient.ID,
		Name:           patient.Name,
		DateOfBirth:    timestamppb.New(*patient.DateBirth),
		Sex:            patient.Sex,
		ProfilePicture: patient.ProfilePicture,
	}, err
}

func (s patientService) UpdatePatientAvatar(stream pb.CreatePatientService_UpdatePatientAvatarServer) error {
	var patientID uint64
	var patient *patient.Patient

	newUUID, err := uuid.NewUUID()
	if err != nil {
		s.log.Error(stream.Context(), "error to generate uuid", log.Body{
			"error": err.Error(),
		})
		return err
	}

	file, err := os.Create("tmp/" + newUUID.String() + ".jpg")
	if err != nil {
		s.log.Error(stream.Context(), "error to generate file by avatar", log.Body{
			"error": err.Error(),
		})
		return err
	}
	defer os.Remove(file.Name())
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Error(stream.Context(), "error to receive avatar chunk", log.Body{
				"error": err.Error(),
			})
			return err
		}
		patientID = chunk.PatientId
		s.log.Info(stream.Context(), "update patient avatar", log.Body{
			"id": chunk.PatientId,
		})

		if patient != nil && patient.ID != chunk.PatientId {
			s.log.Error(stream.Context(), "invalid patient id", log.Body{
				"error": "invalid patient id sequence",
			}
			return fmt.Errorf("invalid patient id")
		}

		patient, err = s.repository.FindPatientByID(stream.Context(), &patientID)
		if err != nil {
			s.log.Error(stream.Context(), "error to get patient", log.Body{
				"error": err,
			})
			return err
		}

		_, err = file.Write(chunk.ImageData)
		if err != nil {
			return err
		}
	}

	if patient.ProfilePicture != "" {
		s.log.Info(stream.Context(), "delete patient avatar", log.Body{
			"id": patientID,
		})

		err := s.storage.Delete(patient.ProfilePicture, "avatar")
		if err != nil {
			s.log.Error(stream.Context(), "error to delete patient avatar", log.Body{
				"error": err,
			})
			return err
		}
	}

	err = s.storage.Save(newUUID.String(), "avatar")
	if err != nil {
		s.log.Error(stream.Context(), "error to save patient avatar", log.Body{
			"error": err.Error(),
		})
		return err
	}

	patient.ProfilePicture = newUUID.String()

	updated, err := s.repository.UpdatePatient(stream.Context(), patient)
	if err != nil {
		s.log.Error(stream.Context(), "error to update patient", log.Body{
			"error": err,
		})
		return err
	}

	stream.SendAndClose(&pb.Patient{
		ID:             updated.ID,
		Name:           updated.Name,
		DateOfBirth:    timestamppb.New(*updated.DateBirth),
		Sex:            updated.Sex,
		ProfilePicture: updated.ProfilePicture,
	})

	return err
}
