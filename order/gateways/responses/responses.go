package responses

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

type genericErrResponse struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

var (
	errInternalServer = genericErrResponse{Type: "err:internal_server_error", Title: "Internal server error"}
	errBadRequest     = genericErrResponse{Type: "err:bad_request", Title: "Bad request"}
	errNotFound       = genericErrResponse{Type: "err:not_found", Title: "Order not found"}
)

type Response struct {
	Payload interface{}
	Status  int
	Error   error
}

func SendJSON(w http.ResponseWriter, payload interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	if payload == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(payload)
}

func ErrorResponse(err error) Response {
	if errors.Is(err, domain.ErrOrderValidation) {
		return BadRequest(err)
	}
	if errors.Is(err, domain.ErrPropertyOwnerIDNotFound) {
		return NotFound(err)
	}

	return InternalServerError(err)
}

func InternalServerError(err error) Response {
	return Response{
		Error:   err,
		Payload: errInternalServer,
		Status:  http.StatusInternalServerError,
	}
}

func NotFound(err error) Response {
	return Response{
		Error:   err,
		Payload: errNotFound,
		Status:  http.StatusNotFound,
	}
}

func BadRequest(err error) Response {
	res := Response{
		Error:   err,
		Payload: errBadRequest,
		Status:  http.StatusBadRequest,
	}

	if errors.Is(err, domain.ErrOrderValidation) {
		newPayload := errBadRequest
		newPayload.Title = err.Error()
		res.Payload = newPayload
	}

	return res
}

func Ok(payload interface{}) Response {
	return Response{
		Error:   nil,
		Payload: payload,
		Status:  http.StatusOK,
	}
}

func Created(payload interface{}) Response {
	return Response{
		Error:   nil,
		Payload: payload,
		Status:  http.StatusCreated,
	}
}

func NoContent() Response {
	return Response{
		Error:   nil,
		Payload: nil,
		Status:  http.StatusNoContent,
	}
}
