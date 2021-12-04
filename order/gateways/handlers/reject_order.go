package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func (b OrderHandler) RejectOrder(r *http.Request) responses.Response {
	var reqBody RejectOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		b.log.Info(err)
		return responses.BadRequest(err)
	}

	err := b.usecase.RejectOrder(r.Context(), domain.RejectOrder{
		PropertyOwnerID: reqBody.PropertyOwnerID,
		OrderID:         reqBody.OrderID,
	})

	if err != nil {
		return responses.ErrorResponse(err)
	}

	return responses.NoContent()
}
