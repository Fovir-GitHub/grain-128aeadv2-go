package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
)

func (h *Handler) handleEncryption(w http.ResponseWriter, r *http.Request) {
	var req model.EncryptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := h.srv.Cipher.Encrypt(&req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusOK, resp); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}
