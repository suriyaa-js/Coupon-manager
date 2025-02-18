package handlers

import (
	"encoding/json"
	"net/http"

	"suriyaa.com/coupon-manager/database"
)

type errMessage struct {
	StatusCode int
	Message    string
}

func returnError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	msg := errorMessage(statusCode, message)
	// msg := errMessage{statusCode, message}
	w.WriteHeader(msg.StatusCode)
	json.NewEncoder(w).Encode(msg)
}

func errorMessage(statusCode int, message string) errMessage {
	switch message {
	case database.Conflict:
		return errMessage{http.StatusConflict, message}
	default:
		return errMessage{statusCode, message}
	}
}
