package handler

import (
	"encoding/json"
	"fmt"
	caregiver "github.com/cuida-me/mvp-backend/internal/application/caregiver/contracts"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"net/http"
)

func CreateCaregiver(useCase caregiver.Create) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		req := &dto.CreateCaregiverRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("failed to decode request: %s", err.Error()),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
					Error:        err,
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, apiErr := useCase.Execute(ctx, req)
		if apiErr != nil {
			response, _ := json.Marshal(apiErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.ErrorStatus)
			w.Write(response)

			return
		}

		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)

		return
	}
}
