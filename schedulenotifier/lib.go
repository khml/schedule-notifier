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

	tasks, err := scheduletask.ReadDefine(taskDefs)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

type Scheduler struct {
	Tasks []scheduletask.ScheduleTask
}

func (s *Scheduler) Exec() {
	currentTime := time.Now()

	needToClean := false
	for _, t := range s.Tasks {
		if t.NeedToNotify(currentTime) {
			_ = Notify(t.Name, "Message body", "")
			t.DoneNotify(currentTime)
		}

		if t.IsAllDone() {
			needToClean = true
		}
	}

	if needToClean {
		var tasks []scheduletask.ScheduleTask
		for _, task := range s.Tasks {
			if task.IsAllDone() {
				continue
			}
			tasks = append(tasks, task)
		}
		s.Tasks = tasks
	}
}

func (s *Scheduler) Run() {
	go func() {
		for range time.Tick(time.Second) {
			s.Exec()
		}
	}()
}

func RunTUI(tasks []scheduletask.ScheduleTask) {
	ui.BuildApp(tasks)
}
