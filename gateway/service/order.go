package service

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// AddOrder handler
func (svc *Service) MakeAddOrderGatewayHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeAddOrderGatewayHandler", "time", time.Now().Local())

	var request AddOrderRequest

	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	c.Request().Header.Set("content-type", "application/json")

	request.OrderID = uuid.New()
	request.CreatedAt = time.Now().Local()
	request.UpdatedAt = time.Now().Local()

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(os.Getenv("ADDORDER"), clientID, clientSecret), bytes.NewBuffer(requests))
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

	//To register the interface of type map[string]interface{}{}
	gob.Register(map[string]interface{}{})

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		level.Error(logger).Log("Error", string(responseBody), "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+string(responseBody))
	}

	session, err := store.Get(c.Request(), "response-session")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	data := response["order_response"]
	content, ok := data.(map[string]interface{})
	if !ok {
		level.Error(logger).Log("Error", "NO response", "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: No response")
	}

	session.Values["response"] = fmt.Sprintf("Member Mail:%s\nOrderID: %s\nBrand Name:%s\nBrand Price:%f\nRam Price:%f\nDVD Price:%f\nNet Price:%f", content["member_email"], content["OrderID"], content["brand_name"], content["brand_price"], content["ram_price"], content["dvd_price"], content["net_price"])
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/send-mail")
}

// MakeGetAllOrderResponseGatewayHandler
func (svc *Service) MakeGetAllOrderResponseGatewayHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeGetAllOrderResponseGatewayHandler", "time", time.Now().Local())

	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(os.Getenv("GETALLORDERDETAILS"), clientID, clientSecret), nil)
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

// GETORDERDBYID HANDLER
func (svc *Service) MakeGetOrderByIDHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeGetOrderByIDHandler", "time", time.Now().Local())

	request := c.FormValue("order_id")
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(os.Getenv("GETORDERDETAILSBYID"), clientID, clientSecret, request), nil)
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

// DeleteOrderByID Handler
func (svc *Service) MakeDeleteOrderByIDGatewayHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeDeleteOrderByIDGatewayHandler", "time", time.Now().Local())
	fmt.Println("HOIIIIII")

	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")
	request := c.FormValue("order_id")

	fmt.Println("ClientID", clientID, clientSecret)

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	fmt.Println(fmt.Sprintf(os.Getenv("DELETEORDER"), clientID, clientSecret, request))

	req, err := http.NewRequest("DELETE", fmt.Sprintf(os.Getenv("DELETEORDER"), clientID, clientSecret, request), nil)
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

// UpdateOrderStatus
func (svc *Service) MakeUpdateOrderStatusGatewayHandler(c echo.Context) error {
	logger := log.With(svc.Logger, "method", "MakeAddOrderGatewayHandler", "time", time.Now().Local())
	var request UpdateOrderStatusRequest

	request.OrderID = c.FormValue("order_id")
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	if clientID == "" || clientSecret == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("ClientID orClientSecret is invalid").Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	c.Request().Header.Set("content-type", "application/json")

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+"error: "+err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(os.Getenv("UPDATESTATUS"), clientID, clientSecret), bytes.NewBuffer(requests))
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
