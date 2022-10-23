package google

import (
	"net/http"
)

func (h *handler) handleGet(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
