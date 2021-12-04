package usecases

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (u useCase) ApproveOrder(ctx context.Context, input domain.ApproveOrder) error {
	order, err := u.storage.GetOrderByID(ctx, input.OrderID, input.PropertyOwnerID)
	if err != nil {
		return err
	}

	if err := order.Approve(); err != nil {
		return err
	}

	if err := u.storage.UpdateOrder(ctx, order); err != nil {
		return err
	}

	err = u.eventPublisher.OrderApproved(ctx, input.OrderID, input.PropertyOwnerID, order.OrderedBy)
	if err != nil {
		return err
	}

	return nil
}
