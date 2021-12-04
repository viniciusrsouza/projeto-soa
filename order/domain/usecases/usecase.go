package usecases

import "github.com/viniciusrsouza/projeto-soa/order/domain"

type useCase struct {
	storage domain.OrderStorage
}

func NewOrderUseCase(s domain.OrderStorage) domain.OrderUseCase {
	return &useCase{
		storage: s,
	}
}
