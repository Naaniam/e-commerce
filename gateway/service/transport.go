package service

import (
	"os"

	"github.com/e-commerce/gateway/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func NewEchoServer(svc Handler) *echo.Echo {
	e := echo.New()

	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}

	e.POST("/v1/e-commerce/admin-signup", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeAdminSignUpGatewayHandler))
	e.POST("/v1/e-commerce/member-signup", svc.MakeMemberSignUpGatewayHandler)

	e.POST("/v1/e-commerce/admin-login", svc.MakeAdminLoginGatewayHandler)
	e.POST("/v1/e-commerce/member-login", svc.MakeMemberLoginGatewayHandler)

	e.GET("/v1/e-commerce/search-members", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeGetAllMembersGatewayHandler))
	e.GET("/v1/e-commerce/search-member/", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeGetMemberByMailHandler))

	e.POST("/v1/e-commerce/add-brand", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeAddBrandGatewayHandler))
	e.GET("/v1/e-commerce/search-brand/", svc.MakeGetBrandByNameHandler)
	e.GET("/v1/e-commerce/search-brands", svc.MakeGetAllBrandsGatewayHandler)
	e.PUT("/v1/e-commerce/update-brand-by-name/", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeUpdateBrandByNameHandler))
	e.DELETE("/v1/e-commerce/delete-brand-by-id/", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeDeleteBrandByIDGatewayHandler))

	e.POST("/v1/e-commerce/add-order", middleware.MemberAuthorize([]byte(os.Getenv("SECRET")), svc.MakeAddOrderGatewayHandler))
	e.GET("/send-mail", middleware.MemberAuthorize([]byte(os.Getenv("SECRET")), SendMailHandler))
	e.GET("/v1/e-commerce/get-order-details", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeGetAllOrderResponseGatewayHandler))
	e.GET("/v1/e-commerce/get-order-details/", middleware.MemberAuthorize([]byte(os.Getenv("SECRET")), svc.MakeGetOrderByIDHandler))
	e.PUT("/v1/e-commerce/update-order-status/", middleware.AdminAuthorize([]byte(os.Getenv("SECRET")), svc.MakeUpdateOrderStatusGatewayHandler))
	e.DELETE("/v1/e-commerce/delete-order/", middleware.MemberAuthorize([]byte(os.Getenv("SECRET")), svc.MakeDeleteOrderByIDGatewayHandler))

	return e
}
