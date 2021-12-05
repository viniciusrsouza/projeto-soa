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

	orderRouter.HandleFunc("/{account_id}", middlewares.Authorize(h.CreateOrder, h.log, h.cfg.AuthService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{account_id}/{status}", middlewares.Authorize(h.ListOrders, h.log, h.cfg.AuthService)).Methods(http.MethodGet)
	orderRouter.HandleFunc("/{account_id}/approve", middlewares.Authorize(h.ApproveOrder, h.log, h.cfg.AuthService)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{account_id}/reject", middlewares.Authorize(h.RejectOrder, h.log, h.cfg.AuthService)).Methods(http.MethodPatch)
	return h
}
