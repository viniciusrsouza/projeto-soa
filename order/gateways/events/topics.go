package events

type Topic string

const (
	ScheduleApprovedTopic Topic = "domain_orderservice_schedule_approved_0"
)

func (t Topic) String() string {
	return string(t)
}
