package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/e-commerce/gateway/service"
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

	defer file.Close()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(io.MultiWriter(os.Stdout, file))
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, "service", "gateway", "time", time.Now().Local(), "caller", log.DefaultCaller)
	}

	svc := service.NewService(logger)

	endPoint := service.NewEchoServer(svc)

	level.Info(logger).Log("msg", "service started", "port", "9000", "time", time.Now().Local())

	defer level.Info(logger).Log("msg", "service ended", "port", "9000", "time", time.Now().Local())

	err = endPoint.Start(":9000")
	level.Error(logger).Log("exit", err, "time", time.Now().Local())

}
