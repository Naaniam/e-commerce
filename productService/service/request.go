package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// AdminSignUpDecodeRequest
func AddBrandDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddBrandRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// GetAllMembersDecodeRequest
func GetAllBrandsDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// GetMemberByIDDecodeRequest
func GetBrandByNameDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetBrandByNameRequest

	request.Name = r.FormValue("brand_name")

	return request, nil
}

// UpdateBrandByNameDecodeRequest
func UpdateBrandByNameDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateBrandByNameRequest

	request.ID = r.FormValue("brand_name")

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// DeleteBrandByIDDecodeRequest
func DeleteBrandByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteBrandByIDRequest

	request.ID = r.FormValue("id")

	return request, nil
}
