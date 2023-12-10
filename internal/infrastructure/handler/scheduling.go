package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	scheduling "github.com/cuida-me/mvp-backend/internal/application/scheduling/contracts"
	"github.com/cuida-me/mvp-backend/pkg/context"
	apiErr "github.com/cuida-me/mvp-backend/pkg/errors"
	"github.com/gorilla/mux"
)

func DoneScheduling(useCase scheduling.DoneScheduling) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		vars := mux.Vars(r)

		schedulingId, err := strconv.ParseUint(vars["schedulingID"], 10, 64)
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

		res, apiErr := useCase.Execute(ctx, &schedulingId)
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

func GetScheduling(useCase scheduling.GetScheduling) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		vars := mux.Vars(r)

		schedulingId, err := strconv.ParseUint(vars["schedulingID"], 10, 64)
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

		res, err := useCase.Execute(ctx, &schedulingId)
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

		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

		return
	}
}

func GetWeekScheduling(useCase scheduling.GetWeekScheduling) http.HandlerFunc {
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
		w.WriteHeader(http.StatusOK)
		w.Write(response)

		return
	}
}

func GetSchedulingReport(useCase scheduling.GetReport) http.HandlerFunc {
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

		report, apiErr := useCase.Execute(ctx, patientID)
		if apiErr != nil {
			response, _ := json.Marshal(apiErr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.ErrorStatus)
			w.Write(response)

			return
		}

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=medication_report.csv")

		writer := csv.NewWriter(w)
		defer writer.Flush()

		for _, row := range report {
			if err := writer.Write(row); err != nil {
				fmt.Println("Failed into write CSV:", err)
				return
			}
		}
	}
}

func ScheduleWeekMedication(job scheduling.ScheduleWeekMedication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		res := job.Execute(ctx)

		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)

		return
	}
}
