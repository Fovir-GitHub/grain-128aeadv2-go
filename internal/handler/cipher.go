package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
)

func (h *Handler) handleCipher(isEncryption bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp any
		var err error

		if isEncryption {
			var req model.EncryptionRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, "invalid request")
				return
			}
			resp, err = h.srv.Cipher.Encrypt(&req)
		} else {
			var req model.DecryptionRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, "invalid request")
				return
			}
			resp, err = h.srv.Cipher.Decrypt(&req)
		}

		if err != nil {
			slog.Error("encrypt/decrypt failed", "err", err)
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := writeJSON(w, http.StatusOK, resp); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
