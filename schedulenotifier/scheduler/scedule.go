package scheduler

import (
	notifytiming "schedule-notifier/schedulenotifier/timing"
	"time"
)

type ScheduleTask struct {
	Name    string
	Time    time.Time
	timings *notifytiming.NotifyTiming
}

func NewSchedule(name string, t time.Time) ScheduleTask {
	timings := notifytiming.NewNotifyTiming(t, -5*time.Minute, -10*time.Minute, -15*time.Minute)
	return ScheduleTask{Name: name, Time: t, timings: &timings}
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
