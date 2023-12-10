package middlewares

import (
	"net/http"

	internal "github.com/cuida-me/mvp-backend/pkg/context"
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/gorilla/mux"
)

func HandleRequestID() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := internal.GetField(r.Context(), internal.RequestIDKey)
			if requestID == "" {
				requestID = internal.GenerateRequestID()
			}

			ctx := internal.CtxWithValues(r.Context(), log.Body{
				internal.RequestIDKey: requestID,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
