package schdulenotifier

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type ScheduleTaskDefine struct {
	Name string `yaml:"name"`
	Time string `yaml:"time"`
}

type ScheduleDefine struct {
	Tasks []ScheduleTaskDefine `yaml:"tasks"`
}

func readYamlFile(filepath string) (ScheduleDefine, error) {
	var schedule ScheduleDefine

	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		// log.Fatalf("Error reading YAML file: %v", err)
		return schedule, err
	}

	if err := yaml.Unmarshal(data, &schedule); err != nil {
		// log.Fatalf("Error unmarshalling YAML: %v", err)
		return schedule, err
	}

	return schedule, nil
}

func ReadYaml(filepath string) ([]ScheduleTask, error) {
	scheduleDef, err := readYamlFile(filepath)
	if err != nil {
		return nil, err
	}

	var tasks []ScheduleTask
	location := time.Now().Location()
	for _, def := range scheduleDef.Tasks {
		taskTime, err := ParseDate(def.Time, location)

		if err != nil {
			return nil, err
		}

		task := ScheduleTask{Name: def.Name, Time: taskTime}
		fmt.Printf("%v\n", task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
