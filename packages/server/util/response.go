package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonResponse, _ := json.Marshal(map[string]string{
		"error": err,
	})
	fmt.Fprint(w, string(jsonResponse))
}

func OKResponse(w http.ResponseWriter, resName string, res interface{}) {
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(map[string]interface{}{
		resName: res,
	})
	fmt.Fprint(w, string(jsonResponse))
}

func NotFoundResponse(w http.ResponseWriter, resName string) {
	w.WriteHeader(http.StatusNotFound)
	jsonResponse, _ := json.Marshal(map[string]interface{}{
		resName: nil,
	})
	fmt.Fprint(w, string(jsonResponse))
}