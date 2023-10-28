package scheduletask

import (
	"schedule-notifier/schedulenotifier/schedulesettings"
	"time"
)

type ScheduleTask struct {
	Name string
	Time time.Time
}

func (s ScheduleTask) IsTime(current time.Time, durationRange time.Duration) bool {
	duration := s.Time.Sub(current)

	if duration < 0 {
		return false
	}

	return duration <= durationRange
}

func ReadSettings(taskDefs []schedulesettings.ScheduleTaskDefine) ([]ScheduleTask, error) {
	var tasks []ScheduleTask

	location := time.Now().Location()
	for _, def := range taskDefs {
		taskTime, err := ParseDate(def.Time, location)

		if err != nil {
			return nil, err
		}

		task := ScheduleTask{Name: def.Name, Time: taskTime}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
