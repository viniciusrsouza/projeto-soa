package responses

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

type GenericErrResponse struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

var (
	errInternalServer      = GenericErrResponse{Type: "err:internal_server_error", Title: "Internal server error"}
	errBadRequest          = GenericErrResponse{Type: "err:bad_request", Title: "Bad request"}
	errNotFound            = GenericErrResponse{Type: "err:not_found", Title: "Order not found"}
	errUnprocessableEntity = GenericErrResponse{Type: "err:unprocessable_entity", Title: "Unprocessable entity"}
	errConflict            = GenericErrResponse{Type: "err:conflict", Title: "Conflict"}
	errUnauthorized        = GenericErrResponse{Type: "err:unauthorized", Title: "Unauthorized"}
	errForbidden           = GenericErrResponse{Type: "err:forbidden", Title: "Forbidden"}
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
	if errors.Is(err, domain.ErrOrderApproveRejected) {
		return UnprocessableEntity(err)
	}
	if errors.Is(err, domain.ErrOrderNotFound) {
		return NotFound(err)
	}
	if errors.Is(err, domain.ErrDuplicatedOrderForSchedule) {
		return Conflict(err)
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
	if errors.Is(err, domain.ErrOrderNotFound) {
		return Response{
			Error: err,
			Payload: GenericErrResponse{
				Type:  errNotFound.Type,
				Title: err.Error(),
			},
			Status: http.StatusNotFound,
		}
	}

	return Response{
		Error:   err,
		Payload: errNotFound,
		Status:  http.StatusNotFound,
	}
}

func UnprocessableEntity(err error) Response {
	payload := errUnprocessableEntity
	payload.Title = err.Error()
	return Response{
		Error:   err,
		Payload: payload,
		Status:  http.StatusUnprocessableEntity,
	}
}

func Conflict(err error) Response {
	payload := errConflict
	payload.Title = err.Error()
	return Response{
		Error:   err,
		Payload: payload,
		Status:  http.StatusConflict,
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

func Unauthorized(err error) Response {
	return Response{
		Error:   err,
		Payload: errUnauthorized,
		Status:  http.StatusUnauthorized,
	}
}

func Forbidden(err error) Response {
	return Response{
		Error:   err,
		Payload: errForbidden,
		Status:  http.StatusForbidden,
	}
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
