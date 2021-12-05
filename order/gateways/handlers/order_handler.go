package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/config"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/middlewares"
)

type OrderHandler struct {
	log     *logrus.Entry
	cfg     config.Config
	usecase domain.OrderUseCase
}

func NewOrderHandler(r *mux.Router, usecase domain.OrderUseCase, cfg config.Config, log *logrus.Entry) OrderHandler {
	h := OrderHandler{
		log:     log,
		usecase: usecase,
		cfg:     cfg,
	}

	orderRouter := r.PathPrefix("/orders").Subrouter()

	// TODO cant buy your own schedule
	orderRouter.HandleFunc("/{account_id}", middlewares.Authorize(h.CreateOrder, h.log, h.cfg.AuthService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{account_id}/{status}", middlewares.Handle(h.ListOrders, h.log)).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{account_id}/approve", middlewares.Handle(h.ApproveOrder, h.log)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{account_id}/reject", middlewares.Handle(h.RejectOrder, h.log)).Methods(http.MethodPatch)
	return h
}
