package ui

import (
	"fmt"
	"schedule-notifier/schedulenotifier/scheduletask"
	"time"

	"github.com/rivo/tview"
)

func formatTime(t time.Time) string {
	return t.Format("2006/01/02 15:04")
}

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Time: 15:04:05"))
}

func updateTime(app *tview.Application, view *tview.TextView) {
	for range time.Tick(time.Second) {
		app.QueueUpdateDraw(func() {
			view.SetText(currentTimeString())
		})
	}
}

func buildTaskList(tasks []scheduletask.ScheduleTask) *tview.TextView {
	list := tview.NewTextView()
	list.SetTitle("Schedule").SetBorder(true)

	txt := ""
	for _, task := range tasks {
		txt += fmt.Sprintf("%v : %v \n", task.Name, formatTime(task.Time))
	}

	list.SetText(txt)
	return list
}

func BuildApp(tasks []scheduletask.ScheduleTask) {
	app := tview.NewApplication()

	clock := tview.NewTextView()
	clock.SetText(currentTimeString())

	taskList := buildTaskList(tasks)

	view := tview.NewFlex().
		AddItem(clock, 0, 1, false).
		AddItem(taskList, 0, 1, false)

	go updateTime(app, clock)
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
