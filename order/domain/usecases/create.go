package usecases

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (u useCase) CreateOrder(ctx context.Context, input domain.Create) (domain.Order, error) {
	order, err := domain.NewOrder(input.OrderedBy, input.PropertyID, input.PropertyOwnerID, input.ScheduleID)
	if err != nil {
		return order, err
	}

	err = u.storage.CreateOrder(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
