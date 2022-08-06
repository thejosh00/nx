package raw

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
	"org/sonatype/nx/util"
)

type RawCreateHostedCommand struct {
	Verbose bool `short:"v" long:"verbose" description:"log verbose debug information"`
}

func (cmd *RawCreateHostedCommand) Execute(args []string) error {
	if !cmd.Verbose {
		util.StopLogging()
	}

	err := createHosted("raw-hosted")
	if err != nil {
		return err
	}
	fmt.Println("Created raw hosted repository")
	return nil
}

type raw struct {
	ContentDisposition string `json:"contentDisposition"`
}

type hostedPayload struct {
	Name    string                  `json:"name"`
	Online  bool                    `json:"online"`
	Raw     raw                     `json:"raw"`
	Storage repomodel.HostedStorage `json:"storage"`
}

func createHosted(name string) error {
	payload := hostedPayload{
		Name:   name,
		Online: true,
		Storage: repomodel.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
		Raw: raw{
			ContentDisposition: "ATTACHMENT",
		},
	}

	return api.Post("v1/repositories/raw/hosted", payload, 201)
}
