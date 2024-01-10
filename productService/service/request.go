package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AdminSignUpDecodeRequest
func AddBrandDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddBrandRequest

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	if !isValidClient(clientID, clientSecret)  && clientID != "admin"{
		return nil, fmt.Errorf("UnAuthorized")
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// GetAllBrandsDecodeRequest
func GetAllBrandsDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// GetBrandByNameDecodeRequest
func GetBrandByNameDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetBrandByNameRequest

	request.Name = r.FormValue("brand_name")

	return request, nil
}

var (
	validClient = map[string]string{
		"admin": "secret",
	}
)

// UpdateBrandByNameDecodeRequest
func UpdateBrandByNameDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateBrandByNameRequest

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	request.ID = r.FormValue("brand_name")

	if !isValidClient(clientID, clientSecret) && clientID != "admin" {
		return nil, fmt.Errorf("UnAuthorized")
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// DeleteBrandByIDDecodeRequest
func DeleteBrandByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteBrandByIDRequest

	request.ID = r.FormValue("id")
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	if !isValidClient(clientID, clientSecret)  && clientID != "admin" {
		return nil, fmt.Errorf("UnAuthorized")
	}

	return request, nil
}

func isValidClient(clientID, clientSecret string) bool {
	if secret, exists := validClient[clientID]; exists && secret == clientSecret {
		return true
	}
	return false
}
