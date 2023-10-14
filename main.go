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
	for {
		currentTime := time.Now()
		fmt.Println("current time:", currentTime)

		for _, t := range tasks {
			if t.IsTime(currentTime, duration) {
				_ = schedulenotifier.Notify(t.Name, "Message body", "")
			}
		}

		time.Sleep(duration)
	}
}
