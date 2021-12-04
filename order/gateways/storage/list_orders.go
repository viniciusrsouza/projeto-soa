package storage

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (s orderStorage) ListOrders(ctx context.Context, propertyOwnerID int, status domain.OrderStatus) ([]domain.Order, error) {
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
			property_owner_id = ($1) and status = ($2)
		order by created_at desc
		`

	rows, err := s.Query(ctx, query, propertyOwnerID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order

		if err := rows.Scan(
			&order.ID,
			&order.Status,
			&order.PropertyID,
			&order.OrderedBy,
			&order.PropertyOwnerID,
			&order.ScheduleID,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
