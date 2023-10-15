package schedulenotifier

import (
	"github.com/gen2brain/beeep"
	"schedule-notifier/schedulenotifier/schedulesettings"
	"schedule-notifier/schedulenotifier/scheduletask"
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
