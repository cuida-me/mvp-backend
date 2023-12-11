package caregiver

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/internal/domain"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
)

type createCaregiverUseCase struct {
	repository caregiver.Repository
	firebase   *auth.Client
	log        log.Provider
	apiErr     apiErr.Provider
}

func NewCreateCaregiverUseCase(
	repository caregiver.Repository,
	firebase *auth.Client,
	log log.Provider,
	apiErr apiErr.Provider,
) *createCaregiverUseCase {
	return &createCaregiverUseCase{
		repository: repository,
		firebase:   firebase,
		log:        log,
		apiErr:     apiErr,
	}
}

func (u createCaregiverUseCase) Execute(ctx context.Context, token string) (*dto.CreateCaregiverResponse, *apiErr.Message) {
	verifiedToken, err := u.firebase.VerifyIDToken(ctx, token)
	if err != nil {
		u.log.Error(ctx, "error to verify token to create user", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.Unauthorized("unauthorized")
	}

	user, err := u.firebase.GetUser(ctx, verifiedToken.UID)
	if err != nil {
		u.log.Error(ctx, "error to get user from firebase", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	exists, err := u.repository.FindCaregiverByEmail(ctx, user.Email)
	if err == nil && exists != nil {
		err := fmt.Errorf("email already exists")
		u.log.Error(ctx, "error to create caregiver", log.Body{
			"email": user.Email,
			"error": err.Error(),
		})
		return nil, u.apiErr.BadRequest("email already exists", err)
	}

	u.log.Info(ctx, "creating caregiver", log.Body{
		"name":  user.DisplayName,
		"email": user.Email,
	})

	caregiver := &caregiver.Caregiver{
		Name:   user.DisplayName,
		Email:  user.Email,
		Status: caregiver.CREATED,
		Uid:    verifiedToken.UID,
	}

	u.resolveCaregiverAvatar(caregiver, &user.PhotoURL)

	created, err := u.repository.CreateCaregiver(ctx, caregiver)
	if err != nil {
		u.log.Error(ctx, "error to create caregiver", log.Body{
			"error": err.Error(),
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	response := dto.CreateCaregiverResponse{}

	response.ToDTO(created)

	return &response, nil
}

func (u createCaregiverUseCase) resolveCaregiverAvatar(c *caregiver.Caregiver, avatar *string) {
	if avatar == nil {
		if c.Sex != nil && *c.Sex == domain.MALE {
			c.Avatar = "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fundraw_Male_avatar_g98d.png?alt=media&token=6460ad65-7be4-4fb4-b694-fc4f67bcff8a"
		} else if c.Sex != nil && *c.Sex == domain.FEMALE {
			c.Avatar = "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fundraw_Female_avatar_efig.png?alt=media&token=dcd87b0c-e54e-44b9-8ca7-14c5fba1ac7e"
		} else {
			c.Avatar = "https://firebasestorage.googleapis.com/v0/b/cuidame-67f4d.appspot.com/o/statics%2Fundraw_Male_avatar_g98d.png?alt=media&token=6460ad65-7be4-4fb4-b694-fc4f67bcff8a"
		}
	} else {
		c.Avatar = *avatar
	}
}
