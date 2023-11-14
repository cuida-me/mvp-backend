package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Baiguoshuai1/shadiaosocketio"
	caregiver "github.com/cuida-me/mvp-backend/internal/application/caregiver/contracts"
	dto "github.com/cuida-me/mvp-backend/internal/application/caregiver/dto"
	"github.com/cuida-me/mvp-backend/pkg/context"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/gorilla/mux"
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

func GetCaregiver(useCase caregiver.GetCaregiver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		values := context.GetCtxValues(ctx)

		caregiverID := values["caregiver_id"].(*uint64)
		if caregiverID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("caregiver not found"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, apiErr := useCase.Execute(ctx, caregiverID)
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

func DeleteCaregiver(useCase caregiver.Delete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		values := context.GetCtxValues(ctx)

		caregiverID := values["caregiver_id"].(*uint64)
		if caregiverID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("caregiver not found"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		apiErr := useCase.Execute(ctx, caregiverID)
		if apiErr != nil {
			response, _ := json.Marshal(apiErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.ErrorStatus)
			w.Write(response)

			return
		}

		w.WriteHeader(http.StatusNoContent)

		return
	}
}

func UpdateCaregiver(useCase caregiver.Update) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		req := &dto.UpdateCaregiverRequest{}
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

		values := context.GetCtxValues(ctx)

		caregiverID := values["caregiver_id"].(*uint64)
		if caregiverID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("caregiver not created"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, apiErr := useCase.Execute(ctx, req, caregiverID)
		if apiErr != nil {
			response, _ := json.Marshal(apiErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.ErrorStatus)
			w.Write(response)

			return
		}

		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

		return
	}
}

func LinkPatientDevice(useCase caregiver.LinkPatientDevice, server *shadiaosocketio.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		vars := mux.Vars(r)
		qrToken := vars["qr_token"]

		values := context.GetCtxValues(ctx)

		caregiverID := values["caregiver_id"].(*uint64)
		if caregiverID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("caregiver not created"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, socketID, apiErr := useCase.Execute(ctx, &qrToken, caregiverID)
		if apiErr != nil {
			response, _ := json.Marshal(apiErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.ErrorStatus)
			w.Write(response)

			return
		}

		channel, err := server.GetChannel(socketID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		err = channel.Emit("finish-login", res)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)

		return
	}
}
