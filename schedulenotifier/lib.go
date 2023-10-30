package schedulenotifier

import (
	"schedule-notifier/schedulenotifier/scheduler"
	"schedule-notifier/schedulenotifier/scheduletask"
	"schedule-notifier/schedulenotifier/settings"
	"schedule-notifier/schedulenotifier/ui"
)

func NewScheduler(pathToSetting string) (*scheduler.Scheduler, error) {
	tasks, err := readSchedule(pathToSetting)
	if err != nil {
		return nil, err
	}

	return &scheduler.Scheduler{Schedules: tasks}, nil
}

func RunTUI(tasks []scheduletask.ScheduleTask) {
	ui.BuildApp(tasks)
}

func readSchedule(filepath string) ([]scheduletask.ScheduleTask, error) {
	taskDefs, err := settings.ReadSettingYaml(filepath)
	if err != nil {
		return nil, err
	}

	tasks, err := scheduletask.ReadDefine(taskDefs)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
