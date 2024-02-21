package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"numbers/numbers"
	adapterHTTP "numbers/numbers/adapter/http"
	infraFile "numbers/numbers/infra/file"
)

var ErrServicePortEmpty = fmt.Errorf("SERVICE_PORT must be provided")
var ErrServicePortInvalid = fmt.Errorf("SERVICE_PORT invalid")
var ErrLogLevelEmpty = fmt.Errorf("SERVICE_LOGLEVEL must be provided")
var ErrLogLevelInvalid = fmt.Errorf("SERVICE_LOGLEVEL invalid")

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

	numbersStorage := infraFile.NewStorage()
	numbersGetterService := numbers.NewGetter(numbersStorage)
	numbersGetterHandler := adapterHTTP.NewGetter(numbersGetterService)
	adapterHTTP.SetRoutes(e, numbersGetterHandler)

	e.Use(middleware.CORS())
	e.Logger.SetLevel(logLevel)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func getPort() (int, error) {
	var port int
	portEnv := os.Getenv("SERVICE_PORT")
	if portEnv == "" {
		return 0, ErrServicePortEmpty
	}
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return 0, fmt.Errorf("%w - %w", ErrServicePortInvalid, err)
	}

	return port, nil
}

func getLogLevel() (log.Lvl, error) {
	logLevel := os.Getenv("SERVICE_LOGLEVEL")
	if logLevel == "" {
		return 0, ErrLogLevelEmpty
	}

	switch strings.ToLower(logLevel) {
	case "debug":
		return log.DEBUG, nil
	case "info":
		return log.INFO, nil
	case "error":
		return log.ERROR, nil
	}

	return 0, fmt.Errorf("%w - %s", ErrLogLevelInvalid, logLevel)
}
