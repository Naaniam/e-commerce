package service

import (
	"context"
	"time"

	"github.com/e-commerce/order/models"
	"github.com/e-commerce/order/repository"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Service interface {
	AddOrderService(ctx context.Context, placeOrder *models.PlaceOrder) (*models.OrderResponse, error)
	GetOrderDetailsByIDService(ctx context.Context, order_id string) (*models.OrderResponse, error)
	GetAllOrderDetailsService(ctx context.Context) (*[]models.OrderResponse, error)
	DeleteOrderDetailsService(ctx context.Context, order_id string) error
	UpdateOrderStatusService(ctx context.Context, order_id, status string) error
}

type service struct {
	Repo   repository.Repository
	Logger log.Logger
}

func NewService(rep *repository.Repository, logger log.Logger) Service {
	return &service{Repo: *rep, Logger: logger}
}

// AddOrderService passes its control to the repository where actual DB operatopns takes place to process the incoming request to add the placed orrder details to placed_order table
func (svc *service) AddOrderService(ctx context.Context, placeOrder *models.PlaceOrder) (*models.OrderResponse, error) {
	logger := log.With(svc.Logger, "method", "AddPlaceOrderService", "time", time.Now().Local())

	orderResponse, err := svc.Repo.CreateOrder(ctx, placeOrder)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Successfully placed order", "time", time.Now().Local())
	return orderResponse, nil
}

// GetOrderDetailsByIDService passes its control to the repository where actual DB operatopns takes place to process the incoming request to get the placed order details based on the given OrderID
func (svc *service) GetOrderDetailsByIDService(ctx context.Context, order_id string) (*models.OrderResponse, error) {
	logger := log.With(svc.Logger, "method", "GetOrderDetailsByIDService", "time", time.Now().Local())
	var orderResponse models.OrderResponse

	err := svc.Repo.GetOrderDetailsByID(ctx, &orderResponse, order_id)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Retrived the order details for id "+order_id, "time", time.Now().Local())
	return &orderResponse, nil
}

// GetAllOrderDetailsService passes its control to the repository where actual DB operatopns takes place to process the incoming request to get all the placed order details
func (svc *service) GetAllOrderDetailsService(ctx context.Context) (*[]models.OrderResponse, error) {
	logger := log.With(svc.Logger, "method", "GetAllOrderDetailsService", "time", time.Now().Local())
	var orderResponse []models.OrderResponse

	err := svc.Repo.GetAllOrderDetails(ctx, &orderResponse)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Retrived all the order details", "time", time.Now().Local())
	return &orderResponse, nil
}

// DeleteOrderDetails
func (svc *service) DeleteOrderDetailsService(ctx context.Context, order_id string) error {
	logger := log.With(svc.Logger, "method", "DeleteOrderDetailsService", "time", time.Now().Local())
	var orderResponse models.OrderResponse

	err := svc.Repo.DeleteOrderByID(ctx, &orderResponse, order_id)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("Message", "Deleted the order of ID "+order_id, "time", time.Now().Local())
	return nil
}

// UpdateOrderStatusService
func (svc *service) UpdateOrderStatusService(ctx context.Context, order_id, status string) error {
	logger := log.With(svc.Logger, "method", "UpdateOrderStatusService", "time", time.Now().Local())
	var orderResponse models.OrderResponse

	err := svc.Repo.UpdateOrderStatus(ctx, &orderResponse, order_id, status)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("Message", "Changed the order status of ID "+order_id+" with "+status, "time", time.Now().Local())
	return nil
}
