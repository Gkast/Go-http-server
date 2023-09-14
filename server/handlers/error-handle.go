package handlers

import (
	"http-server/errors"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	log.Printf("Error: %v", err)

	if appErr, ok := err.(*errors.AppError); ok {
		http.Error(w, appErr.Message, appErr.Code)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
