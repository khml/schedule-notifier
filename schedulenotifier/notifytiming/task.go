package notifyscedule

import "time"

type NotifyTime struct {
	Time time.Time
}

func (n *NotifyTime) HasPassed(current time.Time) bool {
	duration := n.Time.Sub(current)
	return duration < 0
}
