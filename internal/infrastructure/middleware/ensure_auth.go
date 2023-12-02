package middlewares

import (
	"firebase.google.com/go/auth"
	"net/http"
	"strconv"
	"strings"

	"github.com/cuida-me/mvp-backend/internal/domain/caregiver"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	internal "github.com/cuida-me/mvp-backend/pkg/context"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/gorilla/mux"
)

func EnsureAuth(logger log.Provider, caregiverRepo caregiver.Repository, firebase *auth.Client) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !requiresAuth(r) {
				next.ServeHTTP(w, r)
				return
			}

			var userType, userID string
			var err error

			token := r.Header.Get("Authorization")
			userType = r.Header.Get("Type")

			if token == "" || userType == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if userType == "patient" {
				_, userID, err = commons.ValidateJwt(token)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					logger.Error(r.Context(), "error to validate jwt", log.Body{
						"error": err.Error(),
					})
					return
				}

				id, err := strconv.ParseUint(userID, 10, 64)
				if err != nil {
					logger.Error(r.Context(), "error to parse string to int", log.Body{
						"error": err.Error(),
					})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				//patient, err := patientRepo.FindPatientByID(r.Context(), &id)
				//if err != nil || patient == nil {
				//	logger.Error(r.Context(), "error to find patient by id", log.Body{
				//		"error": err.Error(),
				//	})
				//	w.WriteHeader(http.StatusUnauthorized)
				//	return
				//}

				caregiver, err := caregiverRepo.FindCaregiverByPatientID(r.Context(), &id)
				if err != nil || caregiver == nil {
					logger.Error(r.Context(), "error to find caregiver by id", log.Body{
						"error": err.Error(),
					})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				if strings.EqualFold(caregiver.Patient.Status, "cancelled") || strings.EqualFold(caregiver.Status, "cancelled") {
					logger.Error(r.Context(), "patient or caregiver is cancelled", log.Body{})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				ctx := internal.CtxWithValues(r.Context(), log.Body{
					"caregiver_id": &caregiver.ID,
					"patient_id":   &caregiver.PatientID,
					"type":         userType,
				})

				next.ServeHTTP(w, r.WithContext(ctx))

			} else if userType == "caregiver" {
				token, err := firebase.VerifyIDToken(r.Context(), token)
				if err != nil {
					logger.Error(r.Context(), "error to validate token in firebase", log.Body{
						"error": err.Error(),
					})
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				caregiver, err := caregiverRepo.FindCaregiverByUid(r.Context(), token.UID)
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
