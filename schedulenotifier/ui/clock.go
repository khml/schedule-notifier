package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
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

func buildClock(app *tview.Application) *tview.TextView {
	clock := tview.NewTextView()
	clock.SetText(currentTimeString())
	go updateTime(app, clock)

	return clock
}
