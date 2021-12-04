package storage

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (s orderStorage) GetOrderByID(ctx context.Context, orderID, propertyOwnerID int) (domain.Order, error) {
	const query = `
		select
		id,
		status,
		property_id,
		ordered_by,
		property_owner_id,
		schedule_id,
		created_at,
		updated_at
		from
			orders
		where
			id = $1 and property_owner_id = $2;
		`

	var order domain.Order
	if err := s.QueryRow(ctx, query, orderID, propertyOwnerID).Scan(
		&order.ID,
		&order.Status,
		&order.PropertyID,
		&order.OrderedBy,
		&order.PropertyOwnerID,
		&order.ScheduleID,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return domain.Order{}, domain.ErrOrderNotFound
		}
		return domain.Order{}, err
	}

	return order, nil
}
