package middlewares

import (
	"fmt"
	"github.com/cuida-me/mvp-backend/pkg/commons"
	internal "github.com/cuida-me/mvp-backend/pkg/context"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/gorilla/mux"

	"net/http"
)

func EnsureAuth(logger log.Provider) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userType, userID, err := commons.ValidateJwt(token)

			if userID != "" && err == nil {
				//cached := cache.GetMap(r.Context(), userID)
				//
				//if cached["USER_STATUS"] == internal.BlockedStatus || cached["USER_STATUS"] == internal.CancelledStatus {
				//	w.WriteHeader(http.StatusForbidden)
				//	w.Write([]byte("Your user has been blocked or deleted. Please contact our support support@email.com"))
				//	return
				//}

				ctx := internal.CtxWithValues(r.Context(), log.Body{
					"user_id": userID,
					"type":    userType,
				})

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				logger.Info(r.Context(), fmt.Sprintf("Error in hash of user %s Error %v", userID, err), log.Body{})
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}
