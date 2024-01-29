package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func DecodeJSONBody(r *http.Request, target interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	return err
}
func ParseInt(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	return id, err
}
