package service

import (
	"github.com/go-kit/log"
	"github.com/labstack/echo"
)

type Handler interface {
	MakeGetMemberByMailHandler(c echo.Context) error
	MakeGetAllMembersGatewayHandler(c echo.Context) error
	MakeMemberLoginGatewayHandler(c echo.Context) error
	MakeMemberSignUpGatewayHandler(c echo.Context) error
	MakeAddBrandGatewayHandler(c echo.Context) error
	MakeGetAllBrandsGatewayHandler(c echo.Context) error
	MakeGetBrandByNameHandler(c echo.Context) error
	MakeDeleteBrandByIDGatewayHandler(c echo.Context) error
	MakeUpdateBrandByNameHandler(c echo.Context) error
	MakeAdminLoginGatewayHandler(c echo.Context) error
	MakeAdminSignUpGatewayHandler(c echo.Context) error
	MakeAddOrderGatewayHandler(c echo.Context) error
	MakeGetAllOrderResponseGatewayHandler(c echo.Context) error
	MakeGetOrderByIDHandler(c echo.Context) error
	MakeDeleteOrderByIDGatewayHandler(c echo.Context) error
	MakeUpdateOrderStatusGatewayHandler(c echo.Context) error
}

type Service struct {
	Logger log.Logger
}

func NewService(logger log.Logger) Handler {
	return &Service{Logger: logger}
}
