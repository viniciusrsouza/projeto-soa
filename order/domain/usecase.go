package domain

import "context"

type OrderUseCase interface {
	// List(ctx context.Context) ([]ORder, error)
	ListOrders(ctx context.Context, propertyOwnerID int, status OrderStatus) ([]Order, error)
	ApproveOrder(ctx context.Context, input ApproveOrder) error
	RejectOrder(ctx context.Context, input RejectOrder) error
	CreateOrder(ctx context.Context, input Create) (Order, error)
	// Update(ctx context.Context, input Update) error
	// Delete(ctx context.Context, orderID string) error
	// GetByID(ctx context.Context, orderID string) (ORder, error)
}
