package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gitlab.com/mfcekirdek/budget-management-api/model"
	"gitlab.com/mfcekirdek/budget-management-api/service"
	"go.uber.org/zap"
	"net/http"
)

type PaymentHandler interface {
	SavePayment(ctx echo.Context) error
	GetAllPayments(ctx echo.Context) error
}

type paymentHandler struct {
	logger         *zap.Logger
	paymentService service.PaymentService
}

func NewPaymentHandler(e *echo.Echo, logger *zap.Logger, service service.PaymentService) PaymentHandler {
	ph := &paymentHandler{logger: logger, paymentService: service}
	e.POST("/api/v1/payment", ph.SavePayment)
	e.GET("/api/v1/payment", ph.GetAllPayments)
	return ph
}

func (p paymentHandler) SavePayment(ctx echo.Context) error {
	status := http.StatusCreated
	paymentToCreate := model.Payment{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&paymentToCreate)
	if err != nil {
		return err
	}

	payment, err := p.paymentService.SavePayment(ctx.Request().Context(), &paymentToCreate)
	if err != nil {
		status = http.StatusInternalServerError
		response := model.SavePaymentResponse{Error: err.Error(), Data: nil}
		return ctx.JSON(status, response)
	}

	response := model.SavePaymentResponse{Error: "", Data: payment}
	return ctx.JSON(status, response)
}

func (p paymentHandler) GetAllPayments(ctx echo.Context) error {
	status := http.StatusOK
	payments, err := p.paymentService.GetAllPayments(ctx.Request().Context())
	if err != nil {
		status = http.StatusInternalServerError
		response := model.GetAllPaymentsResponse{Error: err.Error(), Data: nil}
		return ctx.JSON(status, response)
	}
	response := model.GetAllPaymentsResponse{Error: "", Data: payments}
	return ctx.JSON(status, response)
}
