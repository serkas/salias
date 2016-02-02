package api

import (
	"net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{
		"name": "salias",
		"version": "0.1.0"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(info); err != nil {
		panic(err)
	}
}
