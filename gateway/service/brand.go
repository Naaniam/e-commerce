package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// AddBrand handler
func (svc *Service) MakeAddBrandGatewayHandler(c echo.Context) error {
	var request AddBrandRequest

	logger := log.With(svc.Logger, "method", "MakeAddBrandGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		fmt.Println("Error Loading.env", err)
	}

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("request", request)

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("requests", string(requests))

	req, err := http.NewRequest("POST", os.Getenv("ADDPRODUCT"), bytes.NewBuffer(requests))
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	var response map[string]interface{}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}

// GETALLPRODUCTS HANDLER
func (svc *Service) MakeGetAllBrandsGatewayHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeGetAllBrandsGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	req, err := http.NewRequest("GET", os.Getenv("GETALLPRODUCTS"), nil)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	fmt.Println("req", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	fmt.Println("responsebody", string(responseBody))

	var response map[string]interface{}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}

// GETPRODUCTBYNAME HANDLER
func (svc *Service) MakeGetBrandByNameHandler(c echo.Context) error {
	request := c.FormValue("name")

	logger := log.With(svc.Logger, "method", "MakeGetBrandByNameHandler", "time", time.Now().Local())

	fmt.Println("Request", request)

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		fmt.Println("Error Loading.env", "error: "+err.Error())
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(os.Getenv("GETPRODUCTBYNAME"), request), nil)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	fmt.Println("Request:", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	var response map[string]interface{}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}
	return c.JSON(http.StatusOK, response)
}

// DeleteProdudctByID Handler
func (svc *Service) MakeDeleteBrandByIDGatewayHandler(c echo.Context) error {
	request := c.FormValue("id")

	logger := log.With(svc.Logger, "method", "MakeDeleteBrandByIDGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	fmt.Println("Values", fmt.Sprintf(os.Getenv("DELETEPRODUCTBYID"), request))

	req, err := http.NewRequest("DELETE", fmt.Sprintf(os.Getenv("DELETEPRODUCTBYID"), request), nil)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	var response map[string]interface{}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}

// UpdateBrandByID Handler
func (svc *Service) MakeUpdateBrandByNameHandler(c echo.Context) error {
	var request UpdateBrandByIDRequest

	logger := log.With(svc.Logger, "method", "MakeUpdateBrandByNameHandler", "time", time.Now().Local())

	request.Name = c.FormValue("brand_name")

	fmt.Println("requestID", request.Name)

	if err := c.Bind(&request.Data); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(os.Getenv("UPDATEPRODUCTBYNAME"), request.Name), bytes.NewBuffer(requests))
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	var response map[string]interface{}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}
