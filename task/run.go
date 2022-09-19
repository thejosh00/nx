package task

import (
	"encoding/json"
	"fmt"
	"org/sonatype/nx/api"
)

type TaskRunCommand struct {
	UseType    bool `short:"t" long:"useType" description:"use type instead of id to locate task"`
	Positional struct {
		Value string `positional-arg-name:"value"`
	} `positional-args:"yes" required:"true"`
}

type TaskItem struct {
	Id string `json:"id"`
}

type Tasks struct {
	Items []TaskItem `json:"items"`
}

func (cmd *TaskRunCommand) Execute(args []string) error {
	id := cmd.Positional.Value

	if cmd.UseType {
		resp, err := listTasks(cmd.Positional.Value)
		if err != nil {
			return err
		}

		var tasks Tasks
		err = json.Unmarshal([]byte(resp), &tasks)
		if err != nil {
			return err
		}

		id = tasks.Items[0].Id
	}

	err := runTask(id)
	if err != nil {
		return err
	}

	fmt.Println("Ran task", cmd.Positional.Value)
	return nil
}

func runTask(taskId string) error {
	url := "v1/tasks/" + taskId + "/run"

	return api.Post(url, nil, 204)
}
