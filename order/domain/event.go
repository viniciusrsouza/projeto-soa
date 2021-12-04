package domain

import "context"

type EventPublisher interface {
	OrderApproved(ctx context.Context, scheduleID, propertyOwnerID, tenantID int) error
}
