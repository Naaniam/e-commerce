package service

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/labstack/echo"
)

func NewEchoServer(svc Service) *echo.Echo {
	e := echo.New()

	addOrderHandler := kithttp.NewServer(AddOrderEndPoint(svc), AddOrderDecodeRequest, EncodeResponse)
	orderDetailsByIdHandler := kithttp.NewServer(GetOrderDetailsByIDEndPoint(svc), GetOrderDetailsByIDDecodeRequest, EncodeResponse)
	allOrderDetailsHandler := kithttp.NewServer(GetAllOrderDetailsEndPoint(svc), GetAllOrderDetailsDecodeRequest, EncodeResponse)
	updateOrderStatus := kithttp.NewServer(UpdateOrderEndPoint(svc), UpdateOrderStatusDecodeRequest, EncodeResponse)
	deleteOrderhandler := kithttp.NewServer(DeleteOrderEndPoint(svc), DeleteOrderByIDDecodeRequest, EncodeResponse)

	e.POST("/e-commerce/v1/admin/add-order", echo.WrapHandler(addOrderHandler))
	e.GET("/e-commerce/v1/member/get-order-details-by-id/", echo.WrapHandler(orderDetailsByIdHandler))
	e.GET("/e-commerce/v1/admin/get-all-order-details", echo.WrapHandler(allOrderDetailsHandler))
	e.PUT("/e-commerce/v1/admin/update-status/", echo.WrapHandler(updateOrderStatus))
	e.DELETE("/e-commerce/v1/member/delete-order/", echo.WrapHandler(deleteOrderhandler))

	return e
}
