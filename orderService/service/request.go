package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func isValidClient(clientID, clientSecret string) bool {
	if secret, exists := validClient[clientID]; exists && secret == clientSecret {
		return true
	}
	return false
}

var (
	validClient = map[string]string{
		"member": "secret1",
		"admin":  "secret",
	}
)

// AddOrderDecodeRequest
func AddOrderDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddOrderRequest

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	fmt.Println("ClientID", clientID)
	fmt.Println("ClientSecret", clientSecret)

	if !isValidClient(clientID, clientSecret) && clientID != "admin" {
		return nil, fmt.Errorf("UnAuthorized")
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

// GetOrderDetailsByIDDecodeRequest
func GetOrderDetailsByIDDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetOrderDetailsByIDRequest

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	request.OrderID = r.FormValue("order_id")

	if !isValidClient(clientID, clientSecret) && clientID != "member" {
		return nil, fmt.Errorf("UnAuthorized")
	}

	return request, nil
}

// GetAllOrderDetailsDecodeRequest
func GetAllOrderDetailsDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	fmt.Println("ClientID", clientID)
	fmt.Println("ClientSecret", clientSecret)

	if !isValidClient(clientID, clientSecret) && clientID != "admin" {
		return nil, fmt.Errorf("UnAuthorized")
	}

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
