package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"schedule-notifier/schedulenotifier/scheduler"
	"schedule-notifier/schedulenotifier/scheduletask"
)

func buildTaskList(title string, tasks []scheduletask.ScheduleTask) *tview.TextView {
	list := tview.NewTextView()
	list.SetTitle(title).SetBorder(true)

	txt := ""
	for _, task := range tasks {
		txt += fmt.Sprintf("%v : %v \n", task.Name, formatTime(task.Time))
	}

	list.SetText(txt)
	return list
}

func BuildApp(s *scheduler.Scheduler) {
	app := tview.NewApplication()

	clock := buildClock(app)
	todayTaskList := buildTaskList("Today Schedule", s.TodayScheduler())
	taskList := buildTaskList("Schedule", s.Schedules)

	view := tview.NewFlex().
		AddItem(clock, 0, 1, false).
		AddItem(todayTaskList, 0, 1, false).
		AddItem(taskList, 0, 1, false)

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
