package handlers

import (
	json "card-validator/internal/utils/json"
	"log"
	"net/http"
)

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s; path: %s; error: %s", r.Method, r.URL.Path, err.Error())

	resp := response{
		Valid: false,
		Error: &errorResponse{
			Code:    http.StatusInternalServerError,
			Message: "server encountered internal problem", // just to be safe
		},
	}

	json.ResponseJSON(w, http.StatusInternalServerError, resp)
}

func badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s; path: %s; error: %s", r.Method, r.URL.Path, err.Error())

	resp := response{
		Valid: false,
		Error: &errorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		},
	}

	json.ResponseJSON(w, http.StatusBadRequest, resp)
}
