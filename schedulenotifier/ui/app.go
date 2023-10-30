package ui

import (
	"fmt"
	"github.com/rivo/tview"
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

func BuildApp(tasks []scheduletask.ScheduleTask) {
	app := tview.NewApplication()

	clock := buildClock(app)
	taskList := buildTaskList("Schedule", tasks)

	view := tview.NewFlex().
		AddItem(clock, 0, 1, false).
		AddItem(taskList, 0, 1, false)

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
