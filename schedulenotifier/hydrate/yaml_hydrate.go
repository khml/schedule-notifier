package hydrate

import (
	"schedule-notifier/schedulenotifier/scheduler"
	"schedule-notifier/schedulenotifier/settings"
	"time"
)

func parseDate(dateString string, location *time.Location) (time.Time, error) {
	layout := "200601021504" // "YYYYMMDDhhmm"
	return time.ParseInLocation(layout, dateString, location)
}

func ReadDefine(scheduleDef settings.ScheduleDefine) ([]scheduler.ScheduleTask, error) {
	var tasks []scheduler.ScheduleTask

	currentTime := time.Now()
	location := currentTime.Location()
	for _, def := range scheduleDef.Tasks {
		taskTime, err := parseDate(def.Time, location)

		if err != nil {
			return nil, err
		}

		task := scheduler.NewSchedule(def.Name, taskTime)
		task.DoneNotify(currentTime) // Turn off notifications for past scheduled times.
		tasks = append(tasks, task)
	}

	return tasks, nil
}
