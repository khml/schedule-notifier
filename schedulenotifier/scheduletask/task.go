package scheduletask

import (
	notifyscedule "schedule-notifier/schedulenotifier/timing"
	"time"
)

type ScheduleTask struct {
	Name    string
	Time    time.Time
	timings notifyscedule.NotifyTiming
}

func (s *ScheduleTask) NeedToNotify(current time.Time) bool {
	return s.timings.NeedToNotify(current)
}

func (s *ScheduleTask) DoneNotify(current time.Time) {
	s.timings.Done(current)
}

func (s *ScheduleTask) IsAllDone() bool {
	return s.timings.IsEmpty()
}
