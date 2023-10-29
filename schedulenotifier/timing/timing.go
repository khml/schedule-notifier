package notifyscedule

import "time"

type NotifyTiming struct {
	tasks []NotifyTime
}

func NewNotifyTiming(base time.Time, durations ...time.Duration) NotifyTiming {
	timings := []NotifyTime{{Time: base}}

	for _, d := range durations {
		t := NotifyTime{Time: base.Add(d)}
		timings = append(timings, t)
	}

	return NotifyTiming{tasks: timings}
}

func (n *NotifyTiming) IsEmpty() bool {
	return len(n.tasks) == 0
}

func (n *NotifyTiming) NeedToNotify(current time.Time) bool {
	if n.IsEmpty() {
		return false
	}

	for _, task := range n.tasks {
		if task.HasPassed(current) {
			// The time has come
			return true
		}
	}

	return false
}

func (n *NotifyTiming) Done(current time.Time) {
	var notPassedTasks []NotifyTime

	for _, task := range n.tasks {
		if !task.HasPassed(current) {
			notPassedTasks = append(notPassedTasks, task)
		}
	}

	n.tasks = notPassedTasks
}
