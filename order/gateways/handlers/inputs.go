package handlers

import "time"

type CreateOrderRequest struct {
	PropertyID      int `json:"property_id"`
	PropertyOwnerID int `json:"property_owner_id"`
	ScheduleID      int `json:"schedule_id"`
}

type ApproveOrderRequest struct {
	OrderID int `json:"order_id"`
}

type RejectOrderRequest struct {
	OrderID int `json:"order_id"`
}

type OrderResponse struct {
	ID              int       `json:"id"`
	OrderedBy       int       `json:"ordered_by"`
	Status          string    `json:"status"`
	PropertyID      int       `json:"property_id"`
	PropertyOwnerID int       `json:"property_owner_id"`
	ScheduleID      int       `json:"schedule_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ListOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}
