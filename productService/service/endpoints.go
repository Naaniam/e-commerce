package service

import (
	"context"
	"time"

	"github.com/e-commerce/product/models"
	"github.com/go-kit/kit/endpoint"
)

// Endpoint calls the service and service calls the repository
// The endpoint will receive a request, convert to the desired format, invoke the service and return the response structure

// Endpoint for AddBrand
type AddBrandRequest struct {
	models.Brand
}

type AdminAddBrandResponse struct {
	Message string       `json:"message,omitempty"`
	Err     string       `json:"err,omitempty"`
	Brand   models.Brand `json:"brand"`
}

// AddBrand
func AddBrandEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddBrandRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		err = s.AddBrandService(ctx, &req.Brand)
		if err != nil {
			return AdminAddBrandResponse{Err: err.Error()}, err
		}

		return AdminAddBrandResponse{Message: "Created the brand " + req.Brand.BrandName + " Successfully!!! ", Brand: req.Brand}, nil
	}
}

//GetAllBrands Endpoint

type GetAllBrandsRequest struct {
	Name string `json:"name"`
}

type GetAllBrandsResponse struct {
	Brands *[]models.Brand `json:"products"`
	Err    string          `json:"err,omitempty"`
}

func MakeGetAllBrandsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		products, err := s.GetAllBrandService(ctx)
		if err != nil {
			return GetAllBrandsResponse{
				Err: err.Error(),
			}, err
		}
		return GetAllBrandsResponse{Brands: products}, nil
	}
}

// GetBrandByName Endpoint

type GetBrandByNameRequest struct {
	Name string `json:"name"`
}

type GetBrandByNameResponse struct {
	Brand *[]models.Brand `json:"brand"`
	Err   string          `json:"err,omitempty"`
}

func MakeGetBrandByNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetBrandByNameRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		brand, err := s.GetBrandByNameService(ctx, req.Name)
		if err != nil {
			return GetAllBrandsResponse{
				Err: err.Error(),
			}, err
		}

		return GetBrandByNameResponse{Brand: brand}, nil
	}
}

// UpdateBrandByName Endpoint
type UpdateBrandByNameRequest struct {
	Data map[string]interface{} `json:"brand"`
	ID   string                 `json:"id"`
}

type UpdateBrandByNameResponse struct {
	Brand *[]models.Brand `json:"brand"`
	Err   string          `json:"err,omitempty"`
}

func MakeUpdateBrandByNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateBrandByNameRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		brand, err := s.UpdateBrandByNameService(ctx, req.Data, req.ID)
		if err != nil {
			return UpdateBrandByNameResponse{
				Err: err.Error(),
			}, err
		}
		return UpdateBrandByNameResponse{Brand: brand}, nil
	}
}

// DeleteBrandByID Endpoint
type DeleteBrandByIDRequest struct {
	ID string `json:"id"`
}

type DeleteBrandByIDResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func MakeDeleteBrandByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		req := request.(DeleteBrandByIDRequest)

		err = s.DeletBrandByIDService(ctx, req.ID)
		if err != nil {
			return DeleteBrandByIDResponse{
				Err: err.Error(),
			}, err
		}
		return DeleteBrandByIDResponse{Message: "Deleted the brand with ID: " + req.ID}, nil
	}
}
