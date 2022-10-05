package docker

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
)

type DockerCreateHostedCommand struct {
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *DockerCreateHostedCommand) Execute(args []string) error {
	name := "docker-hosted"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := createHosted(name)
	if err != nil {
		return err
	}
	fmt.Println("Created docker hosted repository", name)
	return nil
}

type dockerHosted struct {
	V1Enabled      bool `json:"v1Enabled"`
	ForceBasicAuth bool `json:"forceBasicAuth"`
	HttpPort       int  `json:"httpPort"`
}

type hostedPayload struct {
	Name    string                   `json:"name"`
	Online  bool                     `json:"online"`
	Docker  *dockerHosted            `json:"docker,omitempty"`
	Storage *repomodel.HostedStorage `json:"storage,omitempty"`
}

func createHosted(name string) error {
	payload := hostedPayload{
		Name:   name,
		Online: true,
		Storage: &repomodel.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow",
		},
		Docker: &dockerHosted{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HttpPort:       18001,
		},
	}

	return api.Post("v1/repositories/docker/hosted", payload, 201)
}
