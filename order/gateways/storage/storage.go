package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

type orderStorage struct {
	*pgxpool.Pool
}

type OrderStorage interface {
	ListOrders(ctx context.Context, propertyOwnerID int, status domain.OrderStatus) ([]domain.Order, error)
	CreateOrder(ctx context.Context, order *domain.Order) error
}

func NewOrderStorage(db *pgxpool.Pool) OrderStorage {
	return &orderStorage{
		Pool: db,
	}
}
