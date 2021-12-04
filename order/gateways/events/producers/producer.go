package producers

import "github.com/viniciusrsouza/projeto-soa/order/gateways/events"

type Publish interface {
	PublishMessage(topic events.Topic, msg events.Message) error
}

type Producer struct {
	publisher Publish
}

func New(publisher Publish) Producer {
	return Producer{
		publisher: publisher,
	}
}
