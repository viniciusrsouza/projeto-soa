package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/middlewares"
)

type OrderHandler struct {
	log     *logrus.Entry
	usecase domain.OrderUseCase
}

func NewOrderHandler(r *mux.Router, usecase domain.OrderUseCase, log *logrus.Entry) OrderHandler {
	h := OrderHandler{
		log:     log,
		usecase: usecase,
	}

	orderRouter := r.PathPrefix("/orders").Subrouter()

	// orderRouter.HandleFunc("", middlewares.Handle(h.List, h.log)).Methods(http.MethodGet)
	orderRouter.HandleFunc("", middlewares.Handle(h.CreateOrder, h.log)).Methods(http.MethodPost)
	orderRouter.HandleFunc("/{property_owner_id}/{status}", middlewares.Handle(h.ListOrders, h.log)).Methods(http.MethodGet)
	// orderRouter.HandleFunc("/{id}", middlewares.Handle(h.Update, log)).Methods(http.MethodPut)
	// orderRouter.HandleFunc("/{id}", middlewares.Handle(h.Delete, log)).Methods(http.MethodDelete)
	// orderRouter.HandleFunc("/{id}", middlewares.Handle(h.GetByID, log)).Methods(http.MethodGet)
	return h
}
