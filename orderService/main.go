package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/e-commerce/order/drivers"
	"github.com/e-commerce/order/migrations"
	"github.com/e-commerce/order/repository"
	"github.com/e-commerce/order/service"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	//Open or create a log file
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}

	defer file.Close()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(io.MultiWriter(os.Stderr))
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "product", "time", time.Now().Local(), "caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started", "port", "8082", "time", time.Now().Local())

	defer level.Info(logger).Log("msg", "service ended", "port", "8082", "time", time.Now().Local())

	dbConnection := drivers.Connectdatabase()
	migrations.Migrators(dbConnection)

	connection := repository.NewDbConnection(dbConnection, logger)

	e := service.NewEchoServer(service.NewService(&connection, logger))

	err = e.Start(":8082")
	level.Error(logger).Log("exit", err, "time", time.Now().Local())
}
