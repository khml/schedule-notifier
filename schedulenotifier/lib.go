package schedulenotifier

import (
	"github.com/gen2brain/beeep"
	"schedule-notifier/schedulenotifier/schedulesettings"
	"schedule-notifier/schedulenotifier/scheduletask"
	"schedule-notifier/schedulenotifier/ui"
	"time"
)

func Notify(title, message, appIcon string) error {
	return beeep.Notify(title, message, appIcon)
}

func ReadSchedule(filepath string) ([]scheduletask.ScheduleTask, error) {
	taskDefs, err := schedulesettings.ReadSettingYaml(filepath)
	if err != nil {
		return nil, err
	}

	tasks, err := scheduletask.ReadSettings(taskDefs)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

type Scheduler struct {
	Tasks    []scheduletask.ScheduleTask
	Duration time.Duration
}

func (s *Scheduler) Exec() {
	currentTime := time.Now()

	for _, t := range s.Tasks {
		if t.IsTime(currentTime, s.Duration) {
			_ = Notify(t.Name, "Message body", "")
		}
	}
}

func (s *Scheduler) Run() {
	go func() {
		for range time.Tick(time.Second) {
			s.Exec()
		}
	}()
}

func BuildApp(tasks []scheduletask.ScheduleTask) {
	ui.BuildApp(tasks)
}
