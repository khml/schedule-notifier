package main

import (
	"fmt"
	"os"
	"schedule-notifier/schedulenotifier"
)

func main() {
	scheduleDefFile := getScheduleDefFilepathFromCliArgs()

	tasks, err := schedulenotifier.ReadSchedule(scheduleDefFile)
	if err != nil {
		panic(err)
	}

	scheduler := schedulenotifier.Scheduler{Schedules: tasks}
	scheduler.Run()

	schedulenotifier.RunTUI(tasks)
}

func getScheduleDefFilepathFromCliArgs() string {
	args := os.Args

	if len(args) < 2 {
		return "schedule.yaml"
	}

	if args[1] == "-h" {
		fmt.Printf("Usage: %v path/to/schedule.yaml\n", args[0])
		os.Exit(0)
	}

	return args[1]
}
