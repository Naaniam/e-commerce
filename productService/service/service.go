package service

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/e-commerce/product/models"
	"github.com/e-commerce/product/repository"
)

type Service interface {
	AddBrandService(ctx context.Context, brand *models.Brand) error
	GetAllBrandService(ctx context.Context) (*[]models.Brand, error)
	GetBrandByNameService(ctx context.Context, name string) (*[]models.Brand, error)
	UpdateBrandByNameService(ctx context.Context, data map[string]interface{}, id string) (*[]models.Brand, error)
	DeletBrandByIDService(ctx context.Context, id string) error
}

type service struct {
	Repo   repository.Repository
	Logger log.Logger
}

func NewService(rep *repository.Repository, logger log.Logger) Service {
	return &service{Repo: *rep, Logger: logger}
}

// AddBrandService
func (svc *service) AddBrandService(ctx context.Context, brand *models.Brand) error {
	logger := log.With(svc.Logger, "method", "AddBrandService", "time", time.Now().Local())

	err := svc.Repo.CreateBrand(ctx, brand)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("Message", "Successfully added the new brand", "time", time.Now().Local())
	return nil
}

// Get all brands service
func (svc *service) GetAllBrandService(ctx context.Context) (*[]models.Brand, error) {
	var brands []models.Brand

	logger := log.With(svc.Logger, "method", "GetAllBrandService", "time", time.Now().Local())

	err := svc.Repo.GetAllBrands(ctx, &brands)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Successfully retrived all the brands", "time", time.Now().Local())
	return &brands, nil
}

// GetAll
func (svc *service) GetBrandByNameService(ctx context.Context, name string) (*[]models.Brand, error) {
	logger := log.With(svc.Logger, "method", "GetBrandByNameService", "time", time.Now().Local())

	var brand []models.Brand

	err := svc.Repo.GetBrandByName(ctx, &brand, name)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Successfully retrived the brands with name "+name, "time", time.Now().Local())
	return &brand, nil
}

// UpdateBrandByName
func (svc *service) UpdateBrandByNameService(ctx context.Context, data map[string]interface{}, name string) (*[]models.Brand, error) {
	logger := log.With(svc.Logger, "method", "UpdateBrandByNameService", "time", time.Now().Local())
	var brand []models.Brand

	err := svc.Repo.UpdateBrandByName(ctx, data, name)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	err = svc.Repo.GetBrandByName(ctx, &brand, name)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return nil, err
	}

	logger.Log("Message", "Successfully updated the brand details with brand_name: "+name, "time", time.Now().Local())
	return &brand, nil
}

// DeletBrandByID
func (svc *service) DeletBrandByIDService(ctx context.Context, id string) error {
	logger := log.With(svc.Logger, "method", "DeletBrandByIDService", "time", time.Now().Local())

	err := svc.Repo.DeleteBrandByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("Message", "Deleted the brand with id "+id, "time", time.Now().Local())
	return nil
}
