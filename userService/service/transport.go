package service

import (
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/labstack/echo"
)

func NewEchoServer(svc Service) *echo.Echo {
	e := echo.New()

	//AdminHandlers
	adminSignUpHandler := kithttp.NewServer(MakeAdminSignUpEndpoint(svc), AdminSignUpDecodeRequest, EncodeResponse)
	adminLoginHandler := kithttp.NewServer(MakeAdminLoginEndpoint(svc), AdminLoginDecodeRequest, EncodeResponse)
	getAllMembersHandler := kithttp.NewServer(MakeGetAllMembersEndpoint(svc), GetAllMembersDecodeRequest, EncodeResponse)
	getMemberByIDHandler := kithttp.NewServer(MakeGetMemberByIDEndpoint(svc), GetMemberByIDDecodeRequest, EncodeResponse)

	//Member Handlers
	memberSignUpHandler := kithttp.NewServer(MakeMemberSignUpEndpoint(svc), MemberSignUpDecodeRequest, EncodeResponse)
	memberLoginHandler := kithttp.NewServer(MakeMemberLoginEndpoint(svc), MemberLoginDecodeRequest, EncodeResponse)

	//Admin Routes
	e.POST("/e-commerce/v1/admin/signup", echo.WrapHandler(adminSignUpHandler))
	e.POST("/e-commerce/v1/admin/login", echo.WrapHandler(adminLoginHandler))

	e.GET("/e-commerce/v1/admin/get-members", echo.WrapHandler(getAllMembersHandler))
	e.GET("/e-commerce/v1/admin/get-member/", echo.WrapHandler(getMemberByIDHandler))

	//Member Routes to signup, login
	e.POST("/e-commerce/v1/member/login", echo.WrapHandler(memberLoginHandler))
	e.POST("/e-commerce/v1/member/signup", echo.WrapHandler(memberSignUpHandler))

	return e
}
