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

	err := createHosted(name, 18002)
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
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Docker  dockerHosted            `json:"docker"`
	Storage repomodel.HostedStorage `json:"storage"`
}

func createHosted(name string, port int) error {
	payload := hostedPayload{
		Name:   name,
		Online: true,
		Storage: repomodel.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow",
		},
		Docker: dockerHosted{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HttpPort:       port,
		},
	}

	return api.Post("v1/repositories/docker/hosted", payload, 201)
}
