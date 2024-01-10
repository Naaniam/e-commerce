package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AdminSignUpDecodeRequest
func AdminSignUpDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AdminSignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// MemberSignUpDecodeRequest
func MemberSignUpDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request MemberSignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// AdminLoginDecodeRequest
func AdminLoginDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AdminLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil

}

// MemberLoginDecodeRequest
func MemberLoginDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request MemberLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

var (
	validClient = map[string]string{
		"admin": "secret",
	}
)

// GetAllMembersDecodeRequest
func GetAllMembersDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	if !isValidClient(clientID, clientSecret) && clientID != "admin" {
		return nil, fmt.Errorf("UnAuthorized")
	}
	return nil, nil
}

// GetMemberByIDDecodeRequest
func GetMemberByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetMemberByIDRequest

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	request.EmailID = r.FormValue("email")

	if !isValidClient(clientID, clientSecret) && clientID != "admin" {
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
