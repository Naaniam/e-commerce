package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// EncodeResponse function is used to encode the go struct to JSON and write to the response body
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
