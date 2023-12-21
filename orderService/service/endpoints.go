package service

import (
	"context"
	"time"

	"github.com/e-commerce/order/models"
	"github.com/go-kit/kit/endpoint"
)

// Endpoint for AddPlaceOrder

// Requests and Response
type AddOrderRequest struct {
	models.PlaceOrder
}

type AddOrderResponse struct {
	Message       string               `json:"message,omitempty"`
	Err           string               `json:"err,omitempty"`
	OrderResponse models.OrderResponse `json:"order_response"`
}

// AddOrderEndPoint function takes Service as input and returns the endpoint. It processes the incoming request and passes to the AddOrderService and processes the request and returns the response for the order placed
func AddOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddOrderRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		orderResponse, err := s.AddOrderService(ctx, &req.PlaceOrder)
		if err != nil {
			return AddOrderResponse{Err: err.Error()}, err
		}

		if req.PlaceOrder.IsDVDInclude {
			return AddOrderResponse{
				Message:       "Placed the order Successfully",
				OrderResponse: *orderResponse,
			}, nil
		}
		return AddOrderResponse{
			Message:       "Placed the order Successfully",
			OrderResponse: *orderResponse,
		}, nil
	}
}

// GetOrderDetailsByIDEndPoint

// Requests and Response
type GetOrderDetailsByIDRequest struct {
	OrderID string `json:"order_id"`
}

type GetOrderDetailsByIDResponse struct {
	Err           string                `json:"err,omitempty"`
	OrderResponse *models.OrderResponse `json:"order_response"`
}

// Endpoint for GetOrderDetailsByID function takes Service as input and returns the endpoint. It processes the incoming request and passes to the GetOrderDetailsByIDService and processes the request and returns the response as all the placed orders
func GetOrderDetailsByIDEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetOrderDetailsByIDRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		orderResponse, err := s.GetOrderDetailsByIDService(ctx, req.OrderID)
		if err != nil {
			return GetOrderDetailsByIDResponse{Err: err.Error()}, err
		}

		return GetOrderDetailsByIDResponse{
			OrderResponse: orderResponse,
		}, nil
	}
}

// Endpoint for GetOrderDetailsByID
type GetAllOrderDetailsRequest struct {
	OrderID string `json:"order_id"`
}

type GetAllOrderDetailsResponse struct {
	Err           string                  `json:"err,omitempty"`
	OrderResponse *[]models.OrderResponse `json:"order_response"`
}

// GetOrderDetailsByID function created an Endpoint which takes Service interface{} as input and returns the endpoint. It processes the incoming request and passes to the GetOrderDetailsByIDService and processes the request and returns the response as all the placed orders
func GetAllOrderDetailsEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		orderResponse, err := s.GetAllOrderDetailsService(ctx)
		if err != nil {
			return GetOrderDetailsByIDResponse{Err: err.Error()}, err
		}

		return GetAllOrderDetailsResponse{
			OrderResponse: orderResponse,
		}, nil
	}
}

// Endpoint to  Update Order Status
type UpdateOrderStatusRequest struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

type UpdateOrderResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

// UpdateOrderEndPoint function creates an endpoint to update the order status based on the ID
func UpdateOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateOrderStatusRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		err = s.UpdateOrderStatusService(ctx, req.OrderID, req.Status)
		if err != nil {
			return UpdateOrderResponse{Err: err.Error()}, err
		}
		return UpdateOrderResponse{
			Message: "The status of the order with id " + req.OrderID + " is updated with the status: " + req.Status,
		}, nil
	}
}

// Endpoint for DeleteOrder
type DeleteOrderRequest struct {
	OrderID string `json:"order_id"`
}

type DeleteOrderResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

// DeleteOrderEndPoint function creates an endpoint to delete the order based on the ID
func DeleteOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteOrderRequest)
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		err = s.DeleteOrderDetailsService(ctx, req.OrderID)
		if err != nil {
			return DeleteOrderResponse{Err: err.Error()}, err
		}
		return DeleteOrderResponse{
			Message: "Order with id " + req.OrderID + " deleted",
		}, nil
	}
}
