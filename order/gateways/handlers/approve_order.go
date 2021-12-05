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

func (b OrderHandler) ApproveOrder(r *http.Request) responses.Response {
	propertyOwnerID, ok := mux.Vars(r)["account_id"]
	if !ok {
		return responses.BadRequest(fmt.Errorf("missing account_id on route params"))
	}

	ownerID, err := strconv.Atoi(propertyOwnerID)
	if err != nil {
		return responses.BadRequest(fmt.Errorf("invalid account_id"))
	}

	var reqBody ApproveOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		b.log.Info(err)
		return responses.BadRequest(err)
	}

	err = b.usecase.ApproveOrder(r.Context(), domain.ApproveOrder{
		PropertyOwnerID: ownerID,
		OrderID:         reqBody.OrderID,
	})

	if err != nil {
		return responses.ErrorResponse(err)
	}

	return responses.NoContent()
}
