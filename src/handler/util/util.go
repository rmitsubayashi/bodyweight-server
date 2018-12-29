package util

import (
	"net/http"
	"errors"
	"encoding/json"
)

func GetQueryParam(r *http.Request, key string) (string, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return "", errors.New("Incorrect request parameters")
	}

	return keys[0], nil
}

type JSONErrors struct {
	Errors []JSONError `json:"errors"`
}

type JSONError struct {
	Status int `json:"status"`
	Code int `json:"code"`
	Title string `json:"title"`
	Details string `json:"details"`
}

func SendError(w http.ResponseWriter, err error, status int) {
	errorStruct := JSONErrors {
		Errors: []JSONError {
			JSONError {
				Status: status,
				Title: err.Error(),
			},
		},
	}

	json.NewEncoder(w).Encode(errorStruct)

}

func SendData(w http.ResponseWriter, obj interface{}, key string) {
	json.NewEncoder(w).Encode(obj)
}