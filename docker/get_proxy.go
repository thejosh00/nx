package docker

import (
	"fmt"
	"org/sonatype/nx/api"
)

type DockerGetProxyCommand struct {
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *DockerGetProxyCommand) Execute(args []string) error {
	name := "docker-proxy"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := getProxy(name)
	if err != nil {
		return err
	}
	return nil
}

func getProxy(name string) error {
	result, err := api.Get("v1/repositories/docker/proxy/"+name, 200)
	if result != "" {
		fmt.Println(result)
	}
	return err
}
