package application

import (
	"context"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/pb"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type caregiverService struct {
	pb.UnimplementedCaregiverServiceServer
	repository caregiver.Repository
	log        log.Provider
}

func NewCaregiverService(repository caregiver.Repository, log log.Provider) *caregiverService {
	return &caregiverService{
		repository: repository,
		log:        log,
	}
}

func (s caregiverService) Create(ctx context.Context, request *pb.CreateCaregiverRequest) (*pb.Caregiver, error) {
	exists, err := s.repository.FindCaregiverByEmail(ctx, request.Email)
	if err == nil && exists != nil {
		err := fmt.Errorf("email already exists")
		return nil, s.handlerError(ctx, err, "error to create caregiver")
	}

	date, err := commons.ConvertToDate(
		request.Birthdate.Year,
		request.Birthdate.Month,
		request.Birthdate.Day)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to create caregiver")
	}

	s.log.Info(ctx, "creating caregiver", log.Body{
		"name":          request.Name,
		"date_of_birth": date,
		"sex":           request.Sex,
		"email":         request.Email,
	})

	caregiver := &caregiver.Caregiver{
		Name:      request.Name,
		BirthDate: date,
		Sex:       request.Sex,
		Avatar:    request.Avatar,
		Email:     request.Email,
		Status:    caregiver.CREATED,
	}

	s.resolveCaregiverAvatar(caregiver)

	created, err := s.repository.CreateCaregiver(ctx, caregiver)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to create caregiver")
	}

	return created.ToCaregiverDTO(), err
}

func (s caregiverService) FindById(ctx context.Context, request *pb.FindCaregiverByIDRequest) (*pb.CaregiverFull, error) {
	s.log.Info(ctx, "find caregiver", log.Body{
		"id": request.ID,
	})

	caregiver, err := s.repository.FindCaregiverByID(ctx, &request.ID)
	if err != nil {
		return nil, s.handlerError(ctx, err, "error to find caregiver")
	}

	return caregiver.ToCaregiverFullDTO(), err
}

func (s caregiverService) Update(ctx context.Context, request *pb.UpdateCaregiverRequest) (*pb.Blank, error) {
	caregiverSaved, err := s.repository.FindCaregiverByID(ctx, &request.CaregiverID)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to get caregiver")
	}
	if caregiverSaved == nil {
		err := fmt.Errorf("caregiver %v not founded", request.CaregiverID)
		return &pb.Blank{}, s.handlerError(ctx, err, "caregiver not found")
	}

	date, err := commons.ConvertToDate(
		request.Birthdate.Year,
		request.Birthdate.Month,
		request.Birthdate.Day)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to update caregiver")
	}

	caregiver := &caregiver.Caregiver{
		Name:      request.Name,
		BirthDate: date,
		Sex:       request.Sex,
		Avatar:    request.Avatar,
	}

	s.updateCaregiverDiff(caregiverSaved, caregiver)

	s.log.Info(ctx, "updating caregiver", log.Body{
		"id": request.CaregiverID,
	})

	_, err = s.repository.UpdateCaregiver(ctx, caregiverSaved)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to update caregiver")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

func (s caregiverService) Delete(ctx context.Context, request *pb.DeleteCaregiverRequest) (*pb.Blank, error) {
	s.log.Info(ctx, "delete caregiver", log.Body{
		"id": request.ID,
	})

	err := s.repository.DeleteCaregiver(ctx, &request.ID)
	if err != nil {
		return &pb.Blank{}, s.handlerError(ctx, err, "error to delete caregiver")
	}

	return &pb.Blank{}, s.handlerError(ctx, err, "")
}

func (s caregiverService) updateCaregiverDiff(actual, new *caregiver.Caregiver) {
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

func (s caregiverService) resolveCaregiverAvatar(c *caregiver.Caregiver) {
	if c.Avatar == "" {
		if c.Sex == caregiver.MALE {
			// TODO: Implements default image

		} else if c.Sex == caregiver.FEMALE {
			// TODO: Implements default image

		}
	}
}

func (s caregiverService) handlerError(ctx context.Context, err error, title string) error {
	if err != nil {
		s.log.Error(ctx, title, log.Body{
			"error": err.Error(),
		})
		return fmt.Errorf("%s error: %s", title, err.Error())
	}

	return nil
}
