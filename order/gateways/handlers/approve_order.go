package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func (b OrderHandler) ApproveOrder(r *http.Request) responses.Response {
	var reqBody ApproveOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		b.log.Info(err)
		return responses.BadRequest(err)
	}

	err := b.usecase.ApproveOrder(r.Context(), domain.ApproveOrder{
		PropertyOwnerID: reqBody.PropertyOwnerID,
		ScheduleID:      reqBody.ScheduleID,
	})

	if err != nil {
		return responses.ErrorResponse(err)
	}

	return responses.NoContent()
}
