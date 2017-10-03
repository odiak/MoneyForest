package apirouter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"../pgmodels"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg"
)

func NewRouter() http.Handler {
	db := pg.Connect(&pg.Options{
		User:     "kaido",
		Database: "money_forest",
		Addr:     "127.0.0.1:5432",
	})

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Use(handleError)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "hello")
	})
	r.Post("/users", createUser)
	return r
}

type errorResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func handleError(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rv := recover(); rv != nil {
				if logEntry := middleware.GetLogEntry(r); logEntry != nil {
					logEntry.Panic(rv, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rv)
					debug.PrintStack()
				}

				var status int
				var res errorResponse
				switch rv.(type) {
				case pgmodels.ValidationError:
					status = http.StatusBadRequest
					res.ErrorCode = "invalid-parameter"
					res.ErrorMessage = rv.(error).Error()
				default:
					status = http.StatusInternalServerError
					res.ErrorCode = "unknown-error"
					res.ErrorMessage = "unknown error occurred"
				}
				err := renderJsonWithStatus(w, status, res)
				if err != nil {
					panic(err)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func renderJsonWithStatus(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	return e.Encode(v)
}

func renderJson(w http.ResponseWriter, v interface{}) error {
	return renderJsonWithStatus(w, http.StatusOK, v)
}
