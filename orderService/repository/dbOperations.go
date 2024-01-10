package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/e-commerce/order/models"
	"github.com/e-commerce/order/utilities"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"gorm.io/gorm"
)

type DbConnection struct {
	DB     *gorm.DB
	Logger log.Logger
}

type Repository interface {
	CreateOrder(ctx context.Context, placeOrder *models.PlaceOrder) (*models.OrderResponse, error)
	GetOrderDetailsByID(ctx context.Context, orderResponse *models.OrderResponse, orderId string) error
	GetAllOrderDetails(ctx context.Context, orderResponse *[]models.OrderResponse) error
	DeleteOrderByID(ctx context.Context, orderResponse *models.OrderResponse, order_id string) error
	UpdateOrderStatus(ctx context.Context, orderResponse *models.OrderResponse, order_id string, status string) error
}

func NewDbConnection(db *gorm.DB, logger log.Logger) Repository {
	return &DbConnection{DB: db, Logger: log.With(logger, "repo", "gorm", "time", time.Now().Local())}
}

// CreateOrder method
func (db *DbConnection) CreateOrder(ctx context.Context, placeOrder *models.PlaceOrder) (*models.OrderResponse, error) {
	logger := log.With(db.Logger, "repo", "gorm", "method", "CreateOrder", "time", time.Now().Local())
	var member models.Member
	var brand models.Brand
	var ram models.RAMSize
	var orderResponse models.OrderResponse

	err := utilities.ValidateStruct(placeOrder)
	if err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("message: give all the required fields to place the order"), "time", time.Now().Local())
		return nil, err
	}

	if err := db.DB.First(&member, "email=?", placeOrder.MemberMail).Error; err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("Message: The member, "+placeOrder.MemberMail+" is not the member"), "time", time.Now().Local())
		return nil, fmt.Errorf("Message: " + placeOrder.MemberMail + " is not the member")
	}

	if err := db.DB.First(&brand, "brand_name=?", placeOrder.BrandName).Error; err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("Message: The brand, "+placeOrder.BrandName+" is not currently available"), "time", time.Now().Local())
		return nil, fmt.Errorf("message: The brand, " + placeOrder.BrandName + " is not currently available")
	}

	if err := db.DB.First(&ram, "size_in_gb=?", placeOrder.RAMSizeINGB).Error; err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("Message: The brand, "+fmt.Sprintf("%d", placeOrder.RAMSizeINGB)+" is not currently available"), "time", time.Now().Local())
		return nil, fmt.Errorf("Message: The brand, " + fmt.Sprintf("%d", placeOrder.RAMSizeINGB) + " is not currently available")
	}

	if err := db.DB.Debug().Create(&placeOrder).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	if placeOrder.IsDVDInclude {
		orderResponse = models.OrderResponse{
			OrderID:      placeOrder.OrderID,
			IsDVDInclude: placeOrder.IsDVDInclude,
			DVDPrice:     3000,
			OrderStatus:  "Placed Order",
			MemberEmail:  placeOrder.MemberMail,
			BrandName:    placeOrder.BrandName,
			BrandPrice:   brand.BrandPrice,
			RAMPrice:     ram.RamPrice,
			NetPrice:     utilities.NetPrice(brand.BrandPrice, ram.RamPrice, placeOrder.IsDVDInclude),
		}
	}

	orderResponse = models.OrderResponse{
		OrderID:      placeOrder.OrderID,
		IsDVDInclude: placeOrder.IsDVDInclude,
		DVDPrice:     0,
		OrderStatus:  "Placed Order",
		MemberEmail:  placeOrder.MemberMail,
		BrandName:    placeOrder.BrandName,
		BrandPrice:   brand.BrandPrice,
		RAMPrice:     ram.RamPrice,
		NetPrice:     utilities.NetPrice(brand.BrandPrice, ram.RamPrice, placeOrder.IsDVDInclude),
	}

	if err := db.DB.Create(&orderResponse).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	level.Info(logger).Log("Order", fmt.Sprintf("%+v", placeOrder), "time", time.Now().Local())
	return &orderResponse, nil
}

// GetOrderDetails based on the OrderID
func (db *DbConnection) GetOrderDetailsByID(ctx context.Context, orderResponse *models.OrderResponse, orderId string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetOrderDetailsByID", "time", time.Now().Local())

	if err := db.DB.First(&orderResponse, "order_id=?", orderId).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}
	return nil
}

// Get all the order details
func (db *DbConnection) GetAllOrderDetails(ctx context.Context, orderResponse *[]models.OrderResponse) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetAllOrderDetails", "time", time.Now().Local())

	if err := db.DB.Find(&orderResponse).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}
	return nil
}

// Delete the placed order based on OrderID
func (db *DbConnection) DeleteOrderByID(ctx context.Context, orderResponse *models.OrderResponse, order_id string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "DeleteOrderByID", "time", time.Now().Local())

	if err := db.DB.First(&orderResponse, "order_id=?", order_id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Delete(&orderResponse, "order_id=?", order_id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Delete(&models.PlaceOrder{}, "order_id=?", order_id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	return nil
}

// Update the placed order status based on OrderID
func (db *DbConnection) UpdateOrderStatus(ctx context.Context, orderResponse *models.OrderResponse, order_id string, status string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "UpdateOrderStatus", "time", time.Now().Local())

	if err := db.DB.First(&orderResponse, "order_id=?", order_id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Model(&orderResponse).Update("order_status", status).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	return nil
}
