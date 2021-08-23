package util

import (
	"encoding/json"
	"net/http"
)

func ParseJSONBody(r *http.Request) (map[string]string, error) {
	defer r.Body.Close()

	var res map[string]string
	err := json.NewDecoder(r.Body).Decode(&res)

	return res, err
}
