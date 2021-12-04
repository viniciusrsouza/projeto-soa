package usecases

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (s useCase) ListOrders(ctx context.Context, propertyOwnerID int, status domain.OrderStatus) ([]domain.Order, error) {
	if err := status.Validate(); err != nil {
		return nil, err
	}

	if propertyOwnerID <= 0 {
		return nil, domain.ErrPropertyOwnerIDNotFound
	}

	orders, err := s.storage.ListOrders(ctx, propertyOwnerID, status)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
