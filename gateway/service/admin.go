package service

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"bytes"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func (svc *Service) MakeAdminSignUpGatewayHandler(c echo.Context) error {
	var request AdminSignUpRequest

	logger := log.With(svc.Logger, "method", "MakeAdminSignUpGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("POST", os.Getenv("ADMINSIGNUP"), bytes.NewBuffer(requests))
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
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
		return c.JSON(http.StatusBadRequest, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}

// ADMIN LOGIN HANDLER
func (svc *Service) MakeAdminLoginGatewayHandler(c echo.Context) error {
	var request AdminLoginRequest

	logger := log.With(svc.Logger, "method", "MakeAdminLoginGatewayHandler", "time", time.Now().Local())

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	if err := c.Bind(&request); err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusBadRequest, "error: "+err.Error())
	}

	requests, err := json.Marshal(&request)
	if err != nil {
		level.Error(logger).Log("Error", err, "time", time.Now().Local())
		return c.JSON(http.StatusInternalServerError, "error: "+err.Error())
	}

	req, err := http.NewRequest("POST", os.Getenv("ADMINLOGIN"), bytes.NewBuffer(requests))
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
		return c.JSON(http.StatusUnauthorized, "error: "+string(responseBody))
	}

	return c.JSON(http.StatusOK, response)
}
