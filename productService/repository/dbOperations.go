package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/e-commerce/product/models"
	"github.com/e-commerce/product/utilities"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"gorm.io/gorm"
)

type DbConnection struct {
	DB     *gorm.DB
	Logger log.Logger
}

type Repository interface {
	CreateBrand(ctx context.Context, brand *models.Brand) error
	GetAllBrands(ctx context.Context, brands *[]models.Brand) error
	GetBrandByName(ctx context.Context, brand *[]models.Brand, name string) error
	UpdateBrandByName(ctx context.Context, data map[string]interface{}, id string) error
	DeleteBrandByID(ctx context.Context, id string) error
}

func NewDbConnection(db *gorm.DB, logger log.Logger) Repository {
	return &DbConnection{DB: db, Logger: log.With(logger, "repo", "gorm", "time", time.Now().Local())}
}

// create brand
func (db *DbConnection) CreateBrand(ctx context.Context, brand *models.Brand) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "CreateBrand", "time", time.Now().Local())

	err := utilities.ValidateStruct(brand)
	if err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("give all the required fields to add the product"), "time", time.Now().Local())
		return err
	}

	if err := db.DB.Debug().Create(&brand).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("message", "Created the brand")
	level.Info(logger).Log("Brand", fmt.Sprintf("%+v", brand), "time", time.Now().Local())
	return nil
}

// Get all the brands
func (db *DbConnection) GetAllBrands(ctx context.Context, brands *[]models.Brand) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetAllBrands", "time", time.Now().Local())

	if err := db.DB.Debug().Preload("RAMSize").Find(&brands).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if len(*brands) == 0 {
		level.Error(logger).Log("Error", "no brands to display")
		return fmt.Errorf("no brands to display")
	}

	level.Info(logger).Log("Brand", fmt.Sprintf("%+v", brands), "time", time.Now().Local())
	return nil
}

// Get the brand by name
func (db *DbConnection) GetBrandByName(ctx context.Context, brand *[]models.Brand, name string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetBrandByName", "time", time.Now().Local())

	if err := db.DB.Debug().First(&brand, "brand_name = ?", name).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Preload("RAMSize").Find(&brand, "brand_name = ?", name).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	level.Info(logger).Log("Brand", fmt.Sprintf("%+v", brand), "time", time.Now().Local())
	return nil
}

// Update the brand details based on id
func (db *DbConnection) UpdateBrandByName(ctx context.Context, data map[string]interface{}, name string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "UpdateBrandByName")
	var brand models.Brand

	res := db.DB.Debug().First(&models.Brand{}, "brand_name=?", name)
	if res.Error != nil {
		level.Error(logger).Log("Error", res.Error, "time", time.Now().Local())
		return res.Error
	}

	if err := res.Find(&brand, "brand_name=?", name).Error; err != nil {
		level.Error(logger).Log("Error", res.Error, "time", time.Now().Local())
		return res.Error
	}

	for d, value := range data {
		if d == "brand_price" {
			res := db.DB.Model(&models.Brand{}).Where("id=?", brand.ID).Update(d, value)
			if res.Error != nil {
				level.Error(logger).Log("Error", res.Error, "time", time.Now().Local())
				return res.Error
			}

			level.Info(logger).Log("Message", "Successfully updated the "+d+" with value "+fmt.Sprintf("%+v", value), "time", time.Now().Local())
		} else if d == "ram_size" {
			var result []models.RAMSize

			if err := db.DB.Where("brand_id=?", brand.ID).Find(&result).Error; err != nil {
				level.Error(logger).Log("Error", err, "time", time.Now().Local())
				return err
			}

			ids := make([]int, 0)

			for _, i := range result {
				ids = append(ids, int(i.ID))
			}

			resultMap, ok := value.([]interface{})
			if !ok {
				logger.Log("Error", fmt.Errorf("conversion failed"), "time", time.Now().Local())
				return fmt.Errorf("conversion failed")
			}

			if len(resultMap) > len(ids) {
				level.Error(logger).Log("Error", fmt.Errorf("check the brand details for the brand with name %v", brand.ID), "time", time.Now().Local())
				return fmt.Errorf("check the brand details for the brand with name %v", brand.ID)
			}

			for i := 0; i < len(resultMap); i++ {
				key, ok := resultMap[i].(map[string]interface{})
				if !ok {
					level.Error(logger).Log("Error", fmt.Errorf("conversion failed"), "time", time.Now().Local())
					return fmt.Errorf("conversion failed")
				}

				for d, val := range key {
					res := db.DB.Debug().Model(&models.RAMSize{}).Where("brand_id=? AND id=?", brand.ID, ids[i]).Update(d, val)
					if res.Error != nil {
						level.Error(logger).Log("Error", res.Error, "time", time.Now().Local())
						return res.Error
					}

					level.Info(logger).Log("Message", "Successfully updated the "+d+" with value "+fmt.Sprintf("%+v", val), "time", time.Now().Local())
				}
			}
		}
	}

	level.Info(logger).Log("Updated Brand", fmt.Sprintf("%+v", data), "time", time.Now().Local())
	return nil
}

// Delete the brand based on id
func (db *DbConnection) DeleteBrandByID(ctx context.Context, id string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "DeleteBrandByID")
	var brand *models.Brand

	if err := db.DB.First(&brand, id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Delete(&brand).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Delete(&models.RAMSize{}, "brand_id=?", id).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}
	return nil
}
