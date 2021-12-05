package usecases

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (u useCase) RejectOrder(ctx context.Context, input domain.RejectOrder) error {
	order, err := u.storage.GetOrderByID(ctx, input.OrderID, input.PropertyOwnerID)
	if err != nil {
		return err
	}

	if err := order.Reject(); err != nil {
		return err
	}

	if err := u.storage.UpdateOrder(ctx, order); err != nil {
		return err
	}

	return nil
}
