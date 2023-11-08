package caregiver

import (
	"context"
	"fmt"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type createCaregiverUseCase struct {
	repository caregiver.Repository
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewCreateCaregiverUseCase(
	repository caregiver.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
) *createCaregiverUseCase {
	return &createCaregiverUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u createCaregiverUseCase) Execute(ctx context.Context, request *dto.CreateCaregiverRequest) (*dto.CreateCaregiverResponse, *apiErr.Message) {
	exists, err := u.repository.FindCaregiverByEmail(ctx, request.Email)
	if err == nil && exists != nil {
		err := fmt.Errorf("email already exists")
		u.log.Error(ctx, "error to create caregiver", log.Body{
			"email": request.Email,
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest("email already exists", err)
	}

	u.log.Info(ctx, "creating caregiver", log.Body{
		"name":          request.Name,
		"date_of_birth": request.BirthDate,
		"sex":           request.Sex,
		"email":         request.Email,
	})

	caregiver := &caregiver.Caregiver{
		Name:      request.Name,
		BirthDate: request.BirthDate,
		Sex:       domain.Sex(request.Sex),
		Email:     request.Email,
		Status:    caregiver.CREATED,
	}

	u.resolveCaregiverAvatar(caregiver, request.Avatar)

	created, err := u.repository.CreateCaregiver(ctx, caregiver)
	if err != nil {
		u.log.Error(ctx, "error to create caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	// TODO: TEMPORARY FOR TESTS
	jwt, err := commons.NewJwt(fmt.Sprintf("%s_%d", "caregiver", created.ID))
	if err != nil {
		u.log.Error(ctx, "error to create caregiver jwt", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	response := dto.CreateCaregiverResponse{}

	response.ToDTO(created, jwt)

	return &response, nil
}

func (u createCaregiverUseCase) resolveCaregiverAvatar(c *caregiver.Caregiver, avatar *string) {
	if avatar == nil {
		if c.Sex == domain.MALE {
			// TODO: Implements default image

		} else if c.Sex == domain.FEMALE {
			// TODO: Implements default image

		} else {

		}
	} else {
		c.Avatar = *avatar
	}
}
