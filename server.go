package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tylerb/graceful"
	"gitlab.com/mfcekirdek/budget-management-api/config"
	"gitlab.com/mfcekirdek/budget-management-api/database"
	"gitlab.com/mfcekirdek/budget-management-api/handler"
	commonMiddleware "gitlab.com/mfcekirdek/budget-management-api/middleware"
	"gitlab.com/mfcekirdek/budget-management-api/repository"
	"gitlab.com/mfcekirdek/budget-management-api/service"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

type Server struct {
	e      *echo.Echo
	config *config.Config
}

func NewServer(c *config.Config) *Server {
	server := &Server{}

	e := echo.New()
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(commonMiddleware.SetCommonHeaders)

	server.config = c
	server.e = e
	return server
}

func (s *Server) Start() error {
	s.e.Server.Addr = fmt.Sprintf(":%d", s.config.Server.Port)
	logger, err := setupLogger(s.config.IsDebug)
	if err != nil {
		panic(err)
	}

	cluster := database.Setup(&s.config.Couchbase)
	collection := database.GetCollection(s.config.Couchbase.BucketName)

	paymentRepository := repository.NewPaymentRepository(logger, collection, cluster)
	paymentService := service.NewPaymentService(logger, paymentRepository)
	handler.NewPaymentHandler(s.e, logger, paymentService)

	s.e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))
	s.e.GET("/health", s.healthCheck)

	timeout := 10 * time.Second
	return graceful.ListenAndServe(s.e.Server, timeout)
}

func (s *Server) healthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (cv CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func setupLogger(isDebug bool) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if isDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	checkFatalError(err)
	zap.ReplaceGlobals(logger)
	return logger, err
}
