package service

import (
	"context"
	"errors"
	"time"

	"github.com/e-commerce/user/middleware"
	"github.com/e-commerce/user/models"
	"github.com/e-commerce/user/repository"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	MemberSignUpService(ctx context.Context, member *models.Member) error
	AdminSignUpService(ctx context.Context, admin *models.Admin) error
	AdminLoginService(ctx context.Context, email string, password string) (string, error)
	MemberLoginService(ctx context.Context, email string, password string) (string, error)
	GetAllMembersService(ctx context.Context) (*[]models.Member, error)
	GetMemberByIDService(ctx context.Context, email string) (*models.Member, error)
}
type service struct {
	Repo   repository.Repository
	Logger log.Logger
}

func NewService(rep *repository.Repository, logger log.Logger) Service {
	return &service{Repo: *rep, Logger: logger}
}

func (svc *service) MemberSignUpService(ctx context.Context, member *models.Member) error {
	logger := log.With(svc.Logger, "method", "AddProductService", "time", time.Now().Local())

	if member.Email == "" || member.Password == "" {
		level.Error(logger).Log("Error", errors.New("mail or password can't be empty").Error())
		return errors.New("mail or password can't be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(member.Password), 10)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return err
	}

	member.Password = string(hash)

	err = svc.Repo.CreateMember(ctx, member)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return err
	}

	logger.Log("Message", "Successfully Member signed up", "time", time.Now().Local())
	return nil
}

// Admin
func (svc *service) AdminSignUpService(ctx context.Context, admin *models.Admin) error {
	logger := log.With(svc.Logger, "method", "AdminSignUpService", "time", time.Now().Local())

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return err
	}

	admin.Password = string(hash)

	err = svc.Repo.CreateAdmin(ctx, admin)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return err
	}
	logger.Log("Message", "Successfully Admin signed up", "time", time.Now().Local())
	return nil
}

func (svc *service) AdminLoginService(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(svc.Logger, "method", "AdminLoginService", "time", time.Now().Local())

	var admin models.Admin

	if email == "" || password == "" {
		level.Error(logger).Log("Error", errors.New("mail or password can't be empty"))
		return "", errors.New("mail or password can't be empty")
	}

	err := svc.Repo.AdminLoginByMailID(ctx, &admin, email)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}

	//Compare the passwords
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}

	//Generate Token
	token, err := middleware.AdminToken(admin.Email, admin.ID)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}

	logger.Log("Message", "Successfully Admin logged in", "token", token, "time", time.Now().Local())
	return token, nil
}

// MemberLogin service
func (svc *service) MemberLoginService(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(svc.Logger, "method", "MemberLoginService", "time", time.Now().Local())

	var member models.Member

	if email == "" || password == "" {
		level.Error(logger).Log("Error", errors.New("mail or password can't be empty").Error())
		return "", errors.New("mail or password can't be empty")
	}

	err := svc.Repo.MemberLoginByMailID(ctx, &member, email)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}

	//Compare the passwords
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", errors.New("invalid mail or password")
	}

	//Generate Token
	token, err := middleware.MemberToken(member.ID, member.Email)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return "", err
	}

	logger.Log("Message", "Successfully Member logged in", "token", token, "time", time.Now().Local())
	return token, nil
}

// GetAllMembers
func (svc *service) GetAllMembersService(ctx context.Context) (*[]models.Member, error) {
	logger := log.With(svc.Logger, "method", "GetAllMembersService", "time", time.Now().Local())

	var members []models.Member

	err := svc.Repo.GetAllMembers(ctx, &members)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return nil, err
	}

	logger.Log("Message", "Successfully retrived all the members", "time", time.Now().Local())

	return &members, nil
}

// GetMemberByID
func (svc *service) GetMemberByIDService(ctx context.Context, email string) (*models.Member, error) {

	logger := log.With(svc.Logger, "method", "GetMemberByIDService", "time", time.Now().Local())

	var member models.Member

	err := svc.Repo.GetMemberByID(ctx, &member, email)
	if err != nil {
		level.Error(logger).Log("Error", err)
		return nil, err
	}

	logger.Log("Message", "Successfully retrived the member with emailID "+email, "time", time.Now().Local())

	return &member, nil
}
