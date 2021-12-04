package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/viniciusrsouza/projeto-soa/order/gateways/responses"
)

func Handle(handler func(r *http.Request) responses.Response, log *logrus.Entry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.WithContext(r.Context())

		response := handler(r)
		if response.Error != nil {
			log.Error(response.Error)
		}

		err := responses.SendJSON(w, response.Payload, response.Status)
		if err != nil {
			log.Error(err)
		}

		log.Info(response)
	}
}
