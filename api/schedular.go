package api

import (
	"hotstar/pkg/errors"
	"hotstar/pkg/respond"
	"hotstar/store"
	"net/http"
)

func schedularHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	respond.OK(w, store.Slots)
	return nil
}
