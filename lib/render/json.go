package render

import (
	"encoding/json"
	"net/http"
)

// HTTPWriteJSON :
func HTTPWriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	encode(w, code, payload)
}

func encode(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
