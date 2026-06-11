package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/model"
)

func (h *Handler) handleKeyManagement(isWrap bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp any
		var err error

		if isWrap {
			var req model.WrapKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, "invalid request")
				return
			}
			resp, err = h.srv.KeyManagement.WrapKey(&req)
		} else {
			var req model.UnwrapKeyRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, "invalid request")
				return
			}
			resp, err = h.srv.KeyManagement.UnwrapKey(&req)
		}

		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := writeJSON(w, http.StatusOK, resp); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
