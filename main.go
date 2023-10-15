package main

import (
	"fmt"
	"os"
	"schedule-notifier/schedulenotifier"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %v path/to/schedule.yaml\n", args[0])
		return
	}

	tasks, err := schedulenotifier.ReadSchedule(args[1])
	if err != nil {
		panic(err)
	}

	duration := 5 * time.Second
	scheduler := schedulenotifier.Scheduler{Tasks: tasks, Duration: duration}
	scheduler.Run()

	schedulenotifier.BuildApp(tasks)
}
