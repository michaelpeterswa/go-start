package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Healthy bool `json:"healthy"`
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Health{Healthy: true})
}
