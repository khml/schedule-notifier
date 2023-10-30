package scheduletask

import (
	"schedule-notifier/schedulenotifier/settings"
	notifyscedule "schedule-notifier/schedulenotifier/timing"
	"time"
)

func parseDate(dateString string, location *time.Location) (time.Time, error) {
	layout := "200601021504" // "YYYYMMDDhhmm"
	return time.ParseInLocation(layout, dateString, location)
}

func ReadDefine(taskDefs []settings.ScheduleTaskDefine) ([]ScheduleTask, error) {
	var tasks []ScheduleTask

	currentTime := time.Now()
	location := currentTime.Location()
	for _, def := range taskDefs {
		taskTime, err := parseDate(def.Time, location)

		if err != nil {
			return nil, err
		}

		timings := notifyscedule.NewNotifyTiming(taskTime, -5*time.Minute, -10*time.Minute, -15*time.Minute)

		task := ScheduleTask{Name: def.Name, Time: taskTime, timings: &timings}
		task.DoneNotify(currentTime) // Turn off notifications for past scheduled times.
		tasks = append(tasks, task)
	}

	return tasks, nil
}
