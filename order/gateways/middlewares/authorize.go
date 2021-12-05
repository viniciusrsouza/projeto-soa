package middlewares

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/config"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func Authorize(handler func(r *http.Request) responses.Response, log *logrus.Entry, cfg config.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountID, ok := mux.Vars(r)["account_id"]
		if !ok {
			res := responses.BadRequest(fmt.Errorf("missing account_id on route params"))
			responses.SendJSON(w, res.Payload, res.Status)
			return
		}

		accID, err := strconv.Atoi(accountID)
		if err != nil {
			res := responses.BadRequest(fmt.Errorf("invalid account_id"))
			responses.SendJSON(w, res.Payload, res.Status)
			return
		}

		tokenReq, err := authorizeRequestBody(r)
		if err != nil {
			res := responses.Unauthorized(err)
			responses.SendJSON(w, res.Payload, res.Status)
			return
		}

		path := fmt.Sprintf("%s%s", cfg.BaseURL, "/introspection/")
		req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, path, bytes.NewReader(tokenReq))
		if err != nil {
			res := responses.Forbidden(err)
			log.WithError(err).Error("cannot instrospect user with account_id: %s", accountID)
			responses.SendJSON(w, res.Payload, res.Status)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			res := responses.Forbidden(err)
			log.WithError(err).Error("cannot send instrospect request for account_id: %s", accountID)
			responses.SendJSON(w, res.Payload, res.Status)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var authorizeResponse AuthorizeResponse
			if err := json.NewDecoder(resp.Body).Decode(&authorizeResponse); err != nil {
				log.WithError(err).Error("can't decode authorize response")
				res := responses.Forbidden(err)
				responses.SendJSON(w, res.Payload, res.Status)
				return
			}

			if authorizeResponse.ID != accID {
				err := fmt.Errorf(`account_id %d has the wrong token`, accID)
				log.Error(err)
				res := responses.GenericErrResponse{
					Title: err.Error(),
					Type:  "err:forbidden",
				}
				responses.SendJSON(w, res, http.StatusForbidden)
				return
			}
		} else {
			responses.SendJSON(w, resp.Body, resp.StatusCode)
			return
		}

		log = log.WithContext(r.Context())

		response := handler(r)
		if response.Error != nil {
			log.Error(response.Error)
		}

		err = responses.SendJSON(w, response.Payload, response.Status)
		if err != nil {
			log.Error(err)
		}

		log.Info(response)
	}
}

type AuthorizeResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthorizeRequest struct {
	Token string `json:"token"`
}

func authorizeRequestBody(r *http.Request) ([]byte, error) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	if strings.TrimSpace(token) == "" {
		return nil, errors.New("empty token")
	}

	reqBody := AuthorizeRequest{
		Token: token,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf(`encoding body: %w`, err)
	}

	return body, nil
}
