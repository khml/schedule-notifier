package schedulenotifier

import (
	"schedule-notifier/schedulenotifier/scheduletask"
	"schedule-notifier/schedulenotifier/ui"
)

func RunTUI(tasks []scheduletask.ScheduleTask) {
	ui.BuildApp(tasks)
}
