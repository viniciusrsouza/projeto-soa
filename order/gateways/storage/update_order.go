package storage

import (
	"context"

	"github.com/viniciusrsouza/projeto-soa/order/domain"
)

func (s orderStorage) UpdateOrder(ctx context.Context, order domain.Order) error {
	const query = `
		update orders set
			status = $2,
			updated_at = now()
		where id = $1;
		`

	cmd, err := s.Exec(ctx, query, order.ID, order.Status)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() != 1 {
		return domain.ErrOrderNotFound
	}

	return nil
}
