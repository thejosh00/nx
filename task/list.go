package task

import (
	"fmt"
	"org/sonatype/nx/api"
)

type TaskListCommand struct {
	Positional struct {
		Value string `positional-arg-name:"value"`
	} `positional-args:"yes"`
}

func (cmd *TaskListCommand) Execute(args []string) error {
	taskType := cmd.Positional.Value

	resp, err := listTasks(taskType)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

func listTasks(taskType string) (string, error) {
	url := "v1/tasks"
	if taskType != "" {
		url = url + "?type=" + taskType
	}

	return api.Get(url, 200)
}
