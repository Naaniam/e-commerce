package service

import (
	"context"
	"time"

	"github.com/e-commerce/user/models"
	"github.com/go-kit/kit/endpoint"
)

//Endpoint calls the service and service calls the repository
//The endpoint will receive a request, convert to the desired format, invoke the service and return the response structure

// Endpoint for admin signup
type AdminSignUpRequest struct {
	models.Admin
}

type AdminSignUpResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

func MakeAdminSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AdminSignUpRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		err = s.AdminSignUpService(ctx, &req.Admin)
		if err != nil {
			return AdminSignUpResponse{Err: err.Error()}, err
		}

		return AdminSignUpResponse{Message: "created admin successfully!!! Admin Name: " + req.Admin.Name + ", Admin Mail: " + req.Admin.Email}, nil
	}
}

// Endpoint for member signup
type MemberSignUpRequest struct {
	models.Member
}

type MemberSignUpResponse struct {
	Message string `json:"message,omitempty"`
	Err     string `json:"err,omitempty"`
}

func MakeMemberSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MemberSignUpRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		err = s.MemberSignUpService(ctx, &req.Member)
		if err != nil {
			return MemberSignUpResponse{Err: err.Error()}, err
		}

		return MemberSignUpResponse{Message: "created member successfully Email: " + req.Member.Email}, nil
	}
}

// Endpoint for AdminLogin
type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

func MakeAdminLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AdminLoginRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		token, err := s.AdminLoginService(ctx, req.Email, req.Password)
		if err != nil {
			return AdminLoginResponse{Token: "Error creating token", Err: err.Error()}, err
		}

		return AdminLoginResponse{Token: token}, nil
	}
}

// Endpoint for MemberLogin
type MemberLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MemberLoginResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

func MakeMemberLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MemberLoginRequest)

		token, err := s.MemberLoginService(ctx, req.Email, req.Password)
		if err != nil {
			return MemberLoginResponse{Token: "Error creating token", Err: err.Error()}, err
		}

		return MemberLoginResponse{Token: token}, nil
	}
}

// Endpoint for GetAllMembers
type GetAllMembersRequest struct {
	EmailID string `json:"id"`
}

type GetAllMembersResponse struct {
	Members *[]models.Member `json:"members"`
	Err     string           `json:"err,omitempty"`
}

//GetAllMembers Endpoint

func MakeGetAllMembersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		members, err := s.GetAllMembersService(ctx)
		if err != nil {
			return GetAllMembersResponse{
				Err: err.Error(),
			}, err
		}
		return GetAllMembersResponse{Members: members}, nil
	}
}

// Endpoint for GetMembersByID
type GetMemberByIDRequest struct {
	EmailID string `json:"id"`
}

type GetMemberByIDResponse struct {
	Member *models.Member `json:"member"`
	Err    string         `json:"err,omitempty"`
}

// GetMembersByID endpoint creation
func MakeGetMemberByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetMemberByIDRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		member, err := s.GetMemberByIDService(ctx, req.EmailID)
		if err != nil {
			return GetAllMembersResponse{
				Err: err.Error(),
			}, err
		}

		return GetMemberByIDResponse{Member: member}, nil
	}
}
