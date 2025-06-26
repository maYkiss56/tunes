package utilites

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, _ *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func RenderError(w http.ResponseWriter, r *http.Request, status int, message string) {
	RenderJSON(w, r, status, map[string]string{"error": message})
}
