package logger

import (
	"log"
	"mware"
	"net/http"
)

// TODO: print stack?  Print stack in logger vs. print stack in the response?
// Log the response, (or just put the response logger first)

func Recover(logger *log.Logger) mware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func(l *log.Logger) {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					l.Println(err)
				}
			}(logger)

			h.ServeHTTP(w, r)
		})
	}
}
