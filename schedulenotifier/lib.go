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

func RunTUI(s *scheduler.Scheduler) {
	ui.BuildApp(s)
}

func readSchedule(filepath string) ([]scheduletask.ScheduleTask, error) {
	scheduleDef, err := settings.ReadSettingYaml(filepath)
	if err != nil {
		return nil, err
	}

	tasks, err := scheduletask.ReadDefine(scheduleDef)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
