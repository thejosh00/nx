package task

import (
	"fmt"
	"org/sonatype/nx/api"
)

type TaskRunCommand struct {
	Positional struct {
		Value string `positional-arg-name:"value"`
	} `positional-args:"yes" required:"true"`
}

func (cmd *TaskRunCommand) Execute(args []string) error {
	taskId := cmd.Positional.Value

	err := runTask(taskId)
	if err != nil {
		return err
	}

	fmt.Println("Ran task", taskId)
	return nil
}

func runTask(taskId string) error {
	url := "v1/tasks/" + taskId + "/run"

	return api.Post(url, nil, 204)
}
