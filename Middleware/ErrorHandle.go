package middleware

import (
	"log"
	"net/http"
)

// From training if an error occurs log and serve next http requests
func PanicRecoveryHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				log.Printf("Encountered panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}

		}()

		next.ServeHTTP(w, r)

	})

}