package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
)

func (h *Handler) handleWrapKey(w http.ResponseWriter, r *http.Request) {
	var req model.WrapKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}

	k, err := h.srv.KeyManagement.WrapKey(&req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="wrapped.key"`)
	if err := k.Encode(w); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}
