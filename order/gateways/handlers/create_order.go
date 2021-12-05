package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func (b OrderHandler) CreateOrder(r *http.Request) responses.Response {
	orderedByID, ok := mux.Vars(r)["account_id"]
	if !ok {
		return responses.BadRequest(fmt.Errorf("missing account_id on route params"))
	}

	orderOwnerID, err := strconv.Atoi(orderedByID)
	if err != nil {
		return responses.BadRequest(fmt.Errorf("invalid account_id"))
	}

	var reqBody CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		b.log.Info(err)
		return responses.BadRequest(err)
	}

	order, err := b.usecase.CreateOrder(r.Context(), domain.Create{
		PropertyID:      reqBody.PropertyID,
		PropertyOwnerID: reqBody.PropertyOwnerID,
		OrderedBy:       orderOwnerID,
		ScheduleID:      reqBody.ScheduleID,
	})
	if err != nil {
		return responses.ErrorResponse(err)
	}

	response := OrderResponse{
		ID:              order.ID,
		PropertyID:      order.PropertyID,
		Status:          order.Status.String(),
		PropertyOwnerID: order.PropertyOwnerID,
		OrderedBy:       order.OrderedBy,
		ScheduleID:      order.ScheduleID,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	return responses.Created(response)
}
