package storage

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (s orderStorage) CreateOrder(ctx context.Context, order *domain.Order) error {
	const query = `
		insert into orders (
			status,
			property_id,
			ordered_by,
			property_owner_id,
			schedule_id
		) 
		values($1, $2, $3, $4, $5)
		returning id, created_at, updated_at;
	`

	if err := s.QueryRow(ctx, query, order.Status, order.PropertyID, order.OrderedBy, order.PropertyOwnerID, order.ScheduleID).
		Scan(
			&order.ID,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
		return err
	}

	return nil
}
