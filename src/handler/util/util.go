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

type Data struct {
	Data interface{} `json:"data"`
}

func SendData(w http.ResponseWriter, obj interface{}) {
	dataStruct := Data {
		Data: obj,
	}
	json.NewEncoder(w).Encode(dataStruct)
}

func GetData(r *http.Request, objRef interface{}) error {
	if r.Body == nil {
		return errors.New("Empty request body")
	}
	return json.NewDecoder(r.Body).Decode(&objRef)
}
