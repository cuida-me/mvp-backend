package handler

import (
	"encoding/json"
	"fmt"
	patient "github.com/cuida-me/mvp-backend/internal/application/patient/contracts"
	dto "github.com/cuida-me/mvp-backend/internal/application/patient/dto"
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

func NewPatientSession(useCase patient.NewPatientSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("NewPatientSession")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		for {
			var request dto.NewPatientSessionRequest

			err := conn.ReadJSON(&request)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					http.Error(w, "Failed to read message", http.StatusBadRequest)
					return
				}
			}

			res, apiErr := useCase.Execute(r.Context(), &request)
			if apiErr != nil {
				http.Error(w, "Failed to create session", http.StatusInternalServerError)
				return
			}

			if err := conn.WriteJSON(res); err != nil {
				http.Error(w, "Failed to send message", http.StatusInternalServerError)
				return
			}
		}
	}
}
