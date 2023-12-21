package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// AddOrderDecodeRequest
func AddOrderDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// GetOrderDetailsByIDDecodeRequest
func GetOrderDetailsByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetOrderDetailsByIDRequest

	request.OrderID = r.FormValue("order_id")

	return request, nil
}

// GetAllOrderDetailsDecodeRequest
func GetAllOrderDetailsDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// DeleteOrderByIDDecodeRequest
func DeleteOrderByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteOrderRequest

	request.OrderID = r.FormValue("order_id")

	return request, nil
}

// DeleteOrderByIDDecodeRequest
func UpdateOrderStatusDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateOrderStatusRequest

	request.OrderID = r.FormValue("order_id")

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
