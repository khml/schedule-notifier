package scheduler

import (
	"schedule-notifier/schedulenotifier/notify"
	"schedule-notifier/schedulenotifier/scheduletask"
	"time"
)

type Scheduler struct {
	Schedules []scheduletask.ScheduleTask
}

func (s *Scheduler) Exec() {
	currentTime := time.Now()

	needToClean := false
	for _, t := range s.Schedules {
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

func (s *Scheduler) TodayScheduler() (schedules []scheduletask.ScheduleTask) {
	today := time.Now().YearDay()

	for _, schedule := range s.Schedules {
		if schedule.Time.YearDay() == today {
			schedules = append(schedules, schedule)
		}
	}

	return
}

func (s *Scheduler) clean() {
	var tasks []scheduletask.ScheduleTask
	for _, task := range s.Schedules {
		if task.IsAllDone() {
			continue
		}
		tasks = append(tasks, task)
	}
	s.Schedules = tasks
}
