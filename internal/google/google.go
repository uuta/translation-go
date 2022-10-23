package google

import (
	"encoding/json"
	"net/http"
)

func (h *handler) HandleGet(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
