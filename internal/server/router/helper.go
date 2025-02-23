package router

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if _, err = w.Write(j); err != nil {
		return err
	}

	return nil
}
