package producers

import (
	"context"
	"encoding/json"

	"github.com/viniciusrsouza/projeto-soa/order/gateways/events"
)

type OrderApproved struct {
	ScheduleID      int `json:"schedule_id"`
	PropertyOwnerID int `json:"property_owner_id"`
	TenantID        int `json:"tenant_id"`
}

func (p Producer) OrderApproved(ctx context.Context, scheduleID, propertyOwnerID, tenantID int) error {
	payload := OrderApproved{
		ScheduleID:      scheduleID,
		PropertyOwnerID: propertyOwnerID,
		TenantID:        tenantID,
	}

	bytesData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = p.publisher.PublishMessage(events.ScheduleApprovedTopic, bytesData)
	if err != nil {
		return err
	}

	return nil
}
