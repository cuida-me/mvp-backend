package patient

import (
	"context"
	"encoding/json"
	"fmt"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/cache"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/go-redis/redis/v8"
	"time"
)

type newPatientSessionUseCase struct {
	repository patient.Repository
	log        log.Provider
	apiErr     apiErr.Provider
	cache      cache.Provider
}

func NewPatientSessionUseCase(
	repository patient.Repository,
	log log.Provider,
	apiErr apiErr.Provider,
	cache cache.Provider,
) *newPatientSessionUseCase {
	return &newPatientSessionUseCase{
		repository: repository,
		log:        log,
		apiErr:     apiErr,
		cache:      cache,
	}
}

func (u newPatientSessionUseCase) Execute(ctx context.Context, request *dto.NewPatientSessionRequest) (*dto.NewPatientSessionResponse, *apiErr.Message) {
	u.log.Info(ctx, "new patient session", log.Body{
		"ip":        request.Ip,
		"device_id": request.DeviceID,
	})

	var temp TemporaryQRToken

	token := commons.GenerateToken(20)

	cached, err := u.cache.Get(ctx, request.DeviceID)
	if err != nil && err != redis.Nil {
		u.log.Error(ctx, "error on get cache", log.Body{
			"error": err,
		})
		return nil, u.apiErr.InternalServerError(err)
	}

	if cached == "" {
		temp = TemporaryQRToken{
			DeviceID: request.DeviceID,
			Token:    token,
			Ip:       request.Ip,
		}
	} else {
		err := json.Unmarshal([]byte(cached), &temp)
		if err != nil {
			u.log.Error(ctx, "error on unmarshal cached", log.Body{
				"error": err,
			})
			return nil, u.apiErr.InternalServerError(err)
		}
		temp.Token = token
		fmt.Printf("temp: %+v\n", temp)
	}

	dataJSON, err := json.Marshal(temp)

	u.cache.Set(ctx, request.DeviceID, dataJSON, 1*time.Minute)

	return &dto.NewPatientSessionResponse{
		Token: token,
	}, nil
}

type TemporaryQRToken struct {
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
	Ip       string `json:"ip"`
}
