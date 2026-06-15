package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/service"
)

type Handler struct {
	srv *service.Service
}

func New(srv *service.Service) *Handler {
	return &Handler{srv: srv}
}

// Register registers APIs.
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/wrap-key", h.handleKeyManagement(true))
	mux.HandleFunc("/api/unwrap-key", h.handleKeyManagement(false))
	mux.HandleFunc("/api/encrypt", h.handleCipher(true))
	mux.HandleFunc("/api/decrypt", h.handleCipher(false))
}

// writeError returns a error message to frontend.
func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"msg": msg}) //nolint
}

// writeJSON writes JSON format content into response.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
