package schedulenotifier

import (
	"schedule-notifier/schedulenotifier/scheduletask"
	"schedule-notifier/schedulenotifier/settings"
)

func ReadSchedule(filepath string) ([]scheduletask.ScheduleTask, error) {
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
