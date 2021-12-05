package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func (b OrderHandler) ListOrders(r *http.Request) responses.Response {
	propertyOwnerID, ok := mux.Vars(r)["account_id"]
	if !ok {
		return responses.BadRequest(fmt.Errorf("missing account_id on route params"))
	}

	status, ok := mux.Vars(r)["status"]
	if !ok {
		return responses.BadRequest(fmt.Errorf("missing status on route params"))
	}

	ownerID, err := strconv.Atoi(propertyOwnerID)
	if err != nil {
		return responses.BadRequest(fmt.Errorf("invalid account_id"))
	}

	orders, err := b.usecase.ListOrders(r.Context(), ownerID, domain.OrderStatus(status))
	if err != nil {
		return responses.ErrorResponse(err)
	}

	response := make([]OrderResponse, 0, len(orders))

	for _, order := range orders {
		formatedOrder := OrderResponse{
			ID:              order.ID,
			OrderedBy:       order.OrderedBy,
			Status:          order.Status.String(),
			PropertyID:      order.PropertyID,
			PropertyOwnerID: order.PropertyOwnerID,
			ScheduleID:      order.ScheduleID,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       order.UpdatedAt,
		}
		response = append(response, formatedOrder)
	}

	b.log.Info(orders)

	return responses.Ok(response)
}
