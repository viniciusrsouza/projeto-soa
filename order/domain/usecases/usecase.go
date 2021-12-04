package usecases

import "github.com/viniciusrsouza/projeto-soa/order/domain"

type useCase struct {
	storage        domain.OrderStorage
	eventPublisher domain.EventPublisher
}

func NewOrderUseCase(s domain.OrderStorage, e domain.EventPublisher) domain.OrderUseCase {
	return &useCase{
		storage:        s,
		eventPublisher: e,
	}
}
