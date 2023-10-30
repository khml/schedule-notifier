package main

import (
	"fmt"
	"os"
	"schedule-notifier/schedulenotifier"
)

func main() {
	scheduleDefFile := getScheduleDefFilepathFromCliArgs()

	scheduler, err := schedulenotifier.NewScheduler(scheduleDefFile)
	if err != nil {
		panic(err)
	}

	scheduler.Run()

	schedulenotifier.RunTUI(scheduler)
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
