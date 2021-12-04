package handlers

import "time"

type CreateOrderRequest struct {
	OrderedBy       int `json:"ordered_by"`
	PropertyID      int `json:"property_id"`
	PropertyOwnerID int `json:"property_owner_id"`
	ScheduleID      int `json:"schedule_id"`
}

type OrderResponse struct {
	ID              int       `json:"id"`
	OrderedBy       int       `json:"ordered_by"`
	PropertyID      int       `json:"property_id"`
	PropertyOwnerID int       `json:"property_owner_id"`
	ScheduleID      int       `json:"schedule_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ListOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}
