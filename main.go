package main

import (
	"fmt"
	"os"
	"schedule-notifier/schdulenotifier"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %v path/to/schedule.yaml\n", args[0])
		return
	}

	tasks := readTasks(args[1])

	duration := 5 * time.Second
	for true {
		currentTime := time.Now()
		fmt.Println("current time:", currentTime)

		for _, t := range tasks {
			if t.IsTime(currentTime, duration) {
				_ = schdulenotifier.Notify(t.Name, "Message body", "")
			}
		}

		time.Sleep(duration)
	}
}

func readTasks(filepath string) []schdulenotifier.ScheduleTask {
	tasks, err := schdulenotifier.ReadYaml(filepath)
	if err != nil {
		panic(err)
	}
	return tasks
}
