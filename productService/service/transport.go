package service

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/labstack/echo"
)

func NewEchoServer(svc Service) *echo.Echo {
	e := echo.New()

	addBrandsHandler := kithttp.NewServer(AddBrandEndPoint(svc), AddBrandDecodeRequest, EncodeResponse)
	viewAllBrandsHandler := kithttp.NewServer(MakeGetAllBrandsEndpoint(svc), GetAllBrandsDecodeRequest, EncodeResponse)
	viewBrandByNameHandler := kithttp.NewServer(MakeGetBrandByNameEndpoint(svc), GetBrandByNameDecodeRequest, EncodeResponse)
	updateBrandByIDHandler := kithttp.NewServer(MakeUpdateBrandByNameEndpoint(svc), UpdateBrandByNameDecodeRequest, EncodeResponse)
	deleteBrandByIDHandler := kithttp.NewServer(MakeDeleteBrandByIDEndpoint(svc), DeleteBrandByIDDecodeRequest, EncodeResponse)

	e.POST("/e-commerce/v1/admin/add-brand", echo.WrapHandler(addBrandsHandler))
	e.GET("/e-commerce/v1/view-all-brand", echo.WrapHandler(viewAllBrandsHandler))
	e.GET("/e-commerce/v1/view-brand-by-name/", echo.WrapHandler(viewBrandByNameHandler))
	e.PUT("/e-commerce/v1/update-brand-by-name/", echo.WrapHandler(updateBrandByIDHandler))
	e.DELETE("/e-commerce/v1/delete-brand-by-id/", echo.WrapHandler(deleteBrandByIDHandler))

	return e
}
