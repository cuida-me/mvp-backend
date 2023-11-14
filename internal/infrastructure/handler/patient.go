package handler

import (
	"encoding/json"
	"fmt"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/contracts"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
	"github.com/cuida-me/mvp-backend/pkg/context"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func CreatePatient(useCase patient.Create) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		req := &dto.CreatePatientRequest{}
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
					ErrorMessage: fmt.Sprintf("caregiver not found"),
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
		w.WriteHeader(http.StatusCreated)
		w.Write(response)

		return
	}
}

func GetPatient(useCase patient.GetPatient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		values := context.GetCtxValues(ctx)

		patientID := values["patient_id"].(*uint64)
		if patientID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("patient not created"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, apiErr := useCase.Execute(ctx, patientID)
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

func DeletePatient(useCase patient.Delete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		values := context.GetCtxValues(ctx)

		patientID := values["patient_id"].(*uint64)
		if patientID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("patient not created"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		apiErr := useCase.Execute(ctx, patientID)
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

func UpdatePatient(useCase patient.Update) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := r.Context()

		req := &dto.UpdatePatientRequest{}
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

		patientID := values["patient_id"].(*uint64)
		if patientID == nil {
			res, _ := json.Marshal(
				apiErr.Message{
					ErrorMessage: fmt.Sprintf("patient not created"),
					ErrorStatus:  http.StatusBadRequest,
					ErrorCode:    http.StatusText(http.StatusBadRequest),
				})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}

		res, apiErr := useCase.Execute(ctx, req, patientID)
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
