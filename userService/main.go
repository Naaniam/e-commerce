package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/e-commerce/user/drivers"
	"github.com/e-commerce/user/migrations"
	"github.com/e-commerce/user/repository"
	"github.com/e-commerce/user/service"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	// Open or create a log file
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(io.MultiWriter(os.Stdout, file))
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "user", "time", time.Now().Local(), "caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started", "port", "8080", "time", time.Now().Local())

	defer level.Info(logger).Log("msg", "service ended", "port", "8080", "time", time.Now().Local())

	dbConnection := drivers.Connectdatabase()
	migrations.Migrators(dbConnection)

	connection := repository.NewDbConnection(dbConnection, logger)

	e := service.NewEchoServer(service.NewService(&connection, logger))

	e.Start(":8080")
	level.Error(logger).Log("exit", err, "time", time.Now().Local())
}
