package schedulesettings

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func ReadSettingYaml(filepath string) ([]ScheduleTaskDefine, error) {
	scheduleDef, err := readYamlFile(filepath)
	if err != nil {
		return nil, err
	}

	return scheduleDef.Tasks, nil
}
