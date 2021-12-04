package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

	err := s.QueryRow(ctx, query, order.Status, order.PropertyID, order.OrderedBy, order.PropertyOwnerID, order.ScheduleID).
		Scan(
			&order.ID,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return domain.ErrDuplicatedOrderForSchedule
			}
		}
		return err
	}

	return nil
}
