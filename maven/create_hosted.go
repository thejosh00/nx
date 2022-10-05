package maven

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
)

type MavenCreateHostedCommand struct {
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *MavenCreateHostedCommand) Execute(args []string) error {
	name := "maven-hosted"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := createHosted(name)
	if err != nil {
		return err
	}
	fmt.Println("Created maven hosted repository", name)
	return nil
}

type maven struct {
	VersionPolicy      string `json:"versionPolicy"`
	LayoutPolicy       string `json:"layoutPolicy"`
	ContentDisposition string `json:"contentDisposition"`
}

type hostedPayload struct {
	Name    string                   `json:"name"`
	Online  bool                     `json:"online"`
	Maven   *maven                   `json:"maven,omitempty"`
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
		Maven: &maven{
			VersionPolicy:      "MIXED",
			LayoutPolicy:       "STRICT",
			ContentDisposition: "ATTACHMENT",
		},
	}

	return api.Post("v1/repositories/maven/hosted", payload, 201)
}
