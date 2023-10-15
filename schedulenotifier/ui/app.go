package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"schedule-notifier/schedulenotifier/scheduletask"
)

func buildText(s string) *canvas.Text {
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	text := canvas.NewText(s, white)
	return text
}

func BuildApp(tasks []scheduletask.ScheduleTask) {
	a := app.New()

	w := a.NewWindow("Clock")
	w.Resize(fyne.Size{Height: 100, Width: 200})

	clock := buildClock()
	content := container.New(layout.NewVBoxLayout(), clock)

	for _, task := range tasks {
		fmt.Println(task.Name)
		taskNameTxt := buildText(task.Name + " : " + task.Time.Format("01/02 15:04"))
		content.Objects = append(content.Objects, taskNameTxt)
	}

	w.SetContent(content)

	w.ShowAndRun()
}
