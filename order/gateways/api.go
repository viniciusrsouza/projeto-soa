package gateways

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/handlers"
)

type API struct {
	log          *logrus.Entry
	orderUseCase domain.OrderUseCase
}

func NewAPI(useCase domain.OrderUseCase, log *logrus.Entry) API {
	return API{
		orderUseCase: useCase,
		log:          log,
	}
}

func (a API) BuildHandler() http.Handler {
	r := mux.NewRouter()

	routerBasePath := r.PathPrefix("/api/orders-service/v1").Subrouter()

	handlers.NewOrderHandler(routerBasePath, a.orderUseCase, a.log)

	return r
}
