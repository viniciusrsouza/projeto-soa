package domain

type Create struct {
	PropertyID      int
	PropertyOwnerID int
	OrderedBy       int
	ScheduleID      int
}

type ApproveOrder struct {
	PropertyOwnerID int
	ScheduleID      int
}

type RejectOrder struct {
	PropertyOwnerID int
	ScheduleID      int
}
