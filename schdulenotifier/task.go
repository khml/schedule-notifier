package schdulenotifier

import "time"

type ScheduleTask struct {
	Name string
	Time time.Time
}

func (s ScheduleTask) IsTime(current time.Time, durationRange time.Duration) bool {
	duration := s.Time.Sub(current)

	if duration < 0 {
		return false
	}

	return duration <= durationRange
}
