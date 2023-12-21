package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/e-commerce/user/models"
	"github.com/e-commerce/user/utilities"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"gorm.io/gorm"
)

type DbConnection struct {
	DB     *gorm.DB
	Logger log.Logger
}

type Repository interface {
	CreateMember(context.Context, *models.Member) error
	CreateAdmin(context.Context, *models.Admin) error
	AdminLoginByMailID(context.Context, *models.Admin, string) error
	MemberLoginByMailID(context.Context, *models.Member, string) error
	GetAllMembers(ctx context.Context, members *[]models.Member) error
	GetMemberByID(ctx context.Context, member *models.Member, id string) error
}

func NewDbConnection(db *gorm.DB, logger log.Logger) Repository {
	return &DbConnection{DB: db, Logger: log.With(logger, "repo", "gorm", "time", time.Now().Local())}
}

// Method to create admin
func (db *DbConnection) CreateAdmin(ctx context.Context, admin *models.Admin) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "CreateAdmin", "time", time.Now().Local())

	err := utilities.ValidateStruct(admin)
	if err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("give all the required fields to add the product"), "time", time.Now().Local())
		return err
	}

	if err := db.DB.Debug().Create(&admin).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("message", "Created the Admin")
	level.Info(logger).Log("Admin", fmt.Sprintf("%+v", admin), "time", time.Now().Local())

	return nil
}

// Method to create member
func (db *DbConnection) CreateMember(ctx context.Context, member *models.Member) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "CreateMember", "time", time.Now().Local())

	err := utilities.ValidateStruct(member)
	if err != nil {
		level.Error(logger).Log("Error", fmt.Errorf("give all the required fields to add the product"), "time", time.Now().Local())
		return err
	}

	if err := db.DB.Debug().Create(&member).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("message", "Created the Member")
	level.Info(logger).Log("Member", fmt.Sprintf("%+v", member), "time", time.Now().Local())
	return nil
}

// Method to Searchby MailID in Admin table for login purpose
func (db *DbConnection) AdminLoginByMailID(ctx context.Context, admin *models.Admin, email string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "AdminLoginByMailID", "time", time.Now().Local())

	fmt.Println("Admin before in repo:", admin, " email: ", email)

	if err := db.DB.Debug().Find(&admin, "email=?", email).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("message", "admin with mail id "+email+" logged in Successfully")

	return nil
}

// Method to Searchby MailID in Member table for login purpose
func (db *DbConnection) MemberLoginByMailID(ctx context.Context, member *models.Member, email string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "MemberLoginByMailID", "time", time.Now().Local())

	if err := db.DB.Debug().Find(&member, "email=?", email).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	logger.Log("message", "member with mail id "+email+"logged in Successfully")
	return nil
}

// Get all the members
func (db *DbConnection) GetAllMembers(ctx context.Context, members *[]models.Member) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetAllMembers", "time", time.Now().Local())

	if err := db.DB.Debug().Preload("Addresses").Find(&members).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	level.Info(logger).Log("Members", fmt.Sprintf("%+v", members), "time", time.Now().Local())
	return nil
}

func (db *DbConnection) GetMemberByID(ctx context.Context, member *models.Member, email string) error {
	logger := log.With(db.Logger, "repo", "gorm", "method", "GetMemberByID")

	if err := db.DB.First(&member, "email = ?", email).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	if err := db.DB.Preload("Addresses", "member_id IN(?)", member.ID).Find(&member, member.ID).Error; err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return err
	}

	level.Info(logger).Log("Product", fmt.Sprintf("%+v", member), "time", time.Now().Local())
	return nil
}
