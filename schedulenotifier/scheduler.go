package schedulenotifier

import (
	"schedule-notifier/schedulenotifier/notify"
	"schedule-notifier/schedulenotifier/scheduletask"
	"time"
)

type Scheduler struct {
	Tasks []scheduletask.ScheduleTask
}

func (s *Scheduler) Exec() {
	currentTime := time.Now()

	needToClean := false
	for _, t := range s.Tasks {
		if t.NeedToNotify(currentTime) {
			_ = notify.Do(t.Name, "", "")
			t.DoneNotify(currentTime)
		}

		if t.IsAllDone() {
			needToClean = true
		}
	}

	if needToClean {
		s.clean()
	}
}

func (s *Scheduler) Run() {
	go func() {
		for range time.Tick(time.Second) {
			s.Exec()
		}
	}()
}

func (s *Scheduler) clean() {
	var tasks []scheduletask.ScheduleTask
	for _, task := range s.Tasks {
		if task.IsAllDone() {
			continue
		}
		tasks = append(tasks, task)
	}
	s.Tasks = tasks
}