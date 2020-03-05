package api

import (
	"net/http"

	"hotstar/pkg/errors"
	"hotstar/pkg/respond"
	"hotstar/pkg/trace"

	"github.com/go-chi/chi"
)

// API Handler's ---------------------------------------------------------------

// Handler custom api handler help us to handle all the errors in one place
type Handler func(w http.ResponseWriter, r *http.Request) *errors.AppError

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err != nil {
		trace.Log.Infof("StatusCode: %d, Error: %s\n DEBUG: %s\n",
			err.Status, err.Error(), err.Debug)
		respond.Fail(w, err)
	}
}

// Routes - all the registered routes
func Routes(router chi.Router) {
	router.Get("/", IndexHandeler)

	router.Method(http.MethodGet, "/talks", Handler(getAllScheduledTalksHandler))
	router.Method(http.MethodPost, "/talks", Handler(scheduleTalkHandler))
}

// Basic Handler func ---------------------------------------------------------------

// IndexHandeler common index handler for all the service
func IndexHandeler(w http.ResponseWriter, r *http.Request) {
	respond.OK(w, map[string]interface{}{
		"name":    "hotstar-schedular",
		"version": 1,
	})
}
