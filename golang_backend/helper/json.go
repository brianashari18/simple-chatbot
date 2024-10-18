package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequest(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	PanicErr(err)
}

func WriteToResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	PanicErr(err)
}
