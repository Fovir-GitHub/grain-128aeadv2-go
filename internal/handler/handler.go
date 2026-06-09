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

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/wrap-key", h.handleWrapKey)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"msg": msg}) //nolint
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
