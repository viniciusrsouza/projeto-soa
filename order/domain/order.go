package domain

import (
	"errors"
	"fmt"
	"time"
)

type OrderStatus string

func (o OrderStatus) Validate() error {
	if o != PendingOrder && o != ApprovedOrder && o != RejectedOrder {
		return fmt.Errorf("%w: the status %s should've be one of (pending, approved or rejected)", ErrOrderValidation, o)
	}

	return nil
}

func (o OrderStatus) String() string {
	return string(o)
}

const (
	PendingOrder  OrderStatus = "pending"
	ApprovedOrder OrderStatus = "approved"
	RejectedOrder OrderStatus = "rejected"
)

var (
	ErrOrderValidation            = errors.New("invalid order data")
	ErrOrderApproveRejected       = errors.New("can't approve order")
	ErrOrderNotFound              = errors.New("order not found")
	ErrDuplicatedOrderForSchedule = errors.New("can't create more than one order for a schedule")
	ErrPropertyOwnerIDNotFound    = errors.New("no results found for this property_owner_id")
)

type Order struct {
	ID              int
	OrderedBy       int
	ScheduleID      int
	PropertyID      int
	PropertyOwnerID int
	Status          OrderStatus
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewOrder(orderedBy, propertyID, ownerID, scheduleID int) (Order, error) {
	order := Order{
		OrderedBy:       orderedBy,
		PropertyID:      propertyID,
		ScheduleID:      scheduleID,
		PropertyOwnerID: ownerID,
		Status:          PendingOrder,
	}

	if err := order.Validate(); err != nil {
		return Order{}, err
	}

	return order, nil
}

func (s Order) Validate() error {
	if s.PropertyID == 0 {
		return fmt.Errorf("%w: the property_id could not be empty", ErrOrderValidation)
	}
	if s.OrderedBy == 0 {
		return fmt.Errorf("%w: the ordered_by could not be empty", ErrOrderValidation)
	}
	if s.ScheduleID == 0 {
		return fmt.Errorf("%w: the schedule_id could not be empty", ErrOrderValidation)
	}
	return nil
}

func (s *Order) Approve() error {
	if s.Status == ApprovedOrder {
		return fmt.Errorf("%w: order already approved", ErrOrderApproveRejected)
	}
	if s.Status == RejectedOrder {
		return fmt.Errorf("%w: order already rejected", ErrOrderApproveRejected)
	}

	s.Status = ApprovedOrder

	return nil
}

func (s *Order) RejectedOrder() error {
	if s.Status == ApprovedOrder {
		return fmt.Errorf("%w: order already approved", ErrOrderApproveRejected)
	}
	if s.Status == RejectedOrder {
		return fmt.Errorf("%w: order already rejected", ErrOrderApproveRejected)
	}

	s.Status = RejectedOrder

	return nil
}
