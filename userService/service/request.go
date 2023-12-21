package service

import (
	"context"
	"encoding/json"
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

// GetAllMembersDecodeRequest
func GetAllMembersDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// GetMemberByIDDecodeRequest
func GetMemberByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetMemberByIDRequest

	request.EmailID = r.FormValue("email")

	return request, nil
}
