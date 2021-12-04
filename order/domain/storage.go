package domain

import "context"

type OrderStorage interface {
	ListOrders(ctx context.Context, propertyOwnerID int, status OrderStatus) ([]Order, error)
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID, propertyOwnerID int) (Order, error)
	UpdateOrder(ctx context.Context, order Order) error
	// Delete(ctx context.Context, orderID int) error
	// GetByID(ctx context.Context, orderID int) (Order, error)
}
