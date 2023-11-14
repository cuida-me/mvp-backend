package middlewares

import (
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/internal/domain/patient"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	internal "github.com/cuida-me/mvp-backend/pkg/context"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/gorilla/mux"
	"strconv"
	"strings"

	"net/http"
)

func EnsureAuth(logger log.Provider, patientRepo patient.Repository, caregiverRepo caregiver.Repository) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !requiresAuth(r) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")

			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userType, userID, err := commons.ValidateJwt(token)

			if userID != "" && err == nil {
				if userType == "patient" {
					id, err := strconv.ParseUint(userID, 10, 64)
					if err != nil {
						logger.Error(r.Context(), "error to parse string to int", log.Body{
							"error": err.Error(),
						})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					patient, err := patientRepo.FindPatientByID(r.Context(), &id)
					if err != nil || patient == nil {
						logger.Error(r.Context(), "error to find patient by id", log.Body{
							"error": err.Error(),
						})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					caregiver, err := caregiverRepo.FindCaregiverByPatientID(r.Context(), &id)
					if err != nil || caregiver == nil {
						logger.Error(r.Context(), "error to find caregiver by id", log.Body{
							"error": err.Error(),
						})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					if strings.EqualFold(patient.Status, "cancelled") || strings.EqualFold(caregiver.Status, "cancelled") {
						logger.Error(r.Context(), "patient or caregiver is cancelled", log.Body{})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					ctx := internal.CtxWithValues(r.Context(), log.Body{
						"caregiver_id": &caregiver.ID,
						"patient_id":   &patient.ID,
						"type":         userType,
					})

					next.ServeHTTP(w, r.WithContext(ctx))
				} else if userType == "caregiver" {
					id, err := strconv.ParseUint(userID, 10, 64)
					if err != nil {
						logger.Error(r.Context(), "error to parse string to int", log.Body{
							"error": err.Error(),
						})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					caregiver, err := caregiverRepo.FindCaregiverByID(r.Context(), &id)
					if err != nil || caregiver == nil {
						logger.Error(r.Context(), "error to find caregiver by id", log.Body{
							"error": err.Error(),
						})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					if strings.EqualFold(caregiver.Status, "cancelled") {
						logger.Error(r.Context(), "caregiver is cancelled", log.Body{})
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					ctx := internal.CtxWithValues(r.Context(), log.Body{
						"caregiver_id": &caregiver.ID,
						"patient_id":   caregiver.PatientID,
						"type":         userType,
					})

					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					logger.Error(r.Context(), fmt.Sprintf("invalid user type %s id %v", userType, userID), log.Body{})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			} else {
				logger.Error(r.Context(), fmt.Sprintf("Error in hash of user %s Error %v", userID, err), log.Body{})
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}

func requiresAuth(r *http.Request) bool {
	if r.URL.Path == "/caregiver" && r.Method == http.MethodPost {
		return false
	}

	if r.URL.Path == "/ping" && r.Method == http.MethodGet {
		return false
	}
	return true
}
