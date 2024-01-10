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
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(os.Getenv("ADDPRODUCT"), clientID, clientSecret), bytes.NewBuffer(requests))
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

// GETPRODUCTBYNAME HANDLER
func (svc *Service) MakeGetBrandByNameHandler(c echo.Context) error {
	request := c.FormValue("name")

	logger := log.With(svc.Logger, "method", "MakeGetBrandByNameHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(os.Getenv("GETPRODUCTBYNAME"), request), nil)
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

// DeleteProdudctByID Handler
func (svc *Service) MakeDeleteBrandByIDGatewayHandler(c echo.Context) error {
	request := c.FormValue("id")
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	logger := log.With(svc.Logger, "method", "MakeDeleteBrandByIDGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf(os.Getenv("DELETEPRODUCTBYID"), clientID, clientSecret, request), nil)
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
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	logger := log.With(svc.Logger, "method", "MakeUpdateBrandByNameHandler", "time", time.Now().Local())

	request.Name = c.FormValue("brand_name")

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

	req, err := http.NewRequest("PUT", fmt.Sprintf(os.Getenv("UPDATEPRODUCTBYNAME"), clientID, clientSecret, request.Name), bytes.NewBuffer(requests))
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
