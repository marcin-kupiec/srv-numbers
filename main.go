package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/marcin-kupiec/srv-numbers/numbers"
	adapterHTTP "github.com/marcin-kupiec/srv-numbers/numbers/adapter/http"
	infraFile "github.com/marcin-kupiec/srv-numbers/numbers/infra/file"
)

func main() {
	e := echo.New()

	port, err := getPort()
	if err != nil {
		e.Logger.Fatal(err)
	}

	logLevel, err := getLogLevel()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(logLevel)
	e.HTTPErrorHandler = adapterHTTP.HandleError

	// init dependencies
	numbersStorage := infraFile.NewStorage()
	numbersGetterService := numbers.NewService(numbersStorage)
	numbersGetterHandler := adapterHTTP.NewGetHandler(numbersGetterService)

	// setup routes
	adapterHTTP.SetRoutes(e, numbersGetterHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func getPort() (int, error) {
	var port int
	portEnv := os.Getenv("SERVICE_PORT")
	if portEnv == "" {
		return 0, fmt.Errorf("SERVICE_PORT must be provided")
	}
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return 0, fmt.Errorf("SERVICE_PORT invalid - %w", err)
	}

	return port, nil
}

func getLogLevel() (log.Lvl, error) {
	logLevel := os.Getenv("SERVICE_LOGLEVEL")
	if logLevel == "" {
		return 0, fmt.Errorf("SERVICE_LOGLEVEL must be provided")
	}

	switch strings.ToLower(logLevel) {
	case "debug":
		return log.DEBUG, nil
	case "info":
		return log.INFO, nil
	case "error":
		return log.ERROR, nil
	}

	return 0, fmt.Errorf("SERVICE_LOGLEVEL invalid - %s", logLevel)
}
