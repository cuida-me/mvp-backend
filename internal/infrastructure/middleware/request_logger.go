package middlewares

import (
	"github.com/cuida-me/mvp-backend/pkg/log"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestLogger(logger log.Provider) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			body, _ := ioutil.ReadAll(r.Body)

			logger.Info(r.Context(), "request-completed", log.Body{
				"route":        r.Method + " " + r.URL.Path,
				"request_body": body,
				"duration_ms":  time.Since(startTime).Milliseconds(),
				"IP":           r.RemoteAddr,
			})
		})
	}
}
