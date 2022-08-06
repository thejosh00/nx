package blobstore

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/util"
)

type BlobstoreCreateFileCommand struct {
	Verbose    bool `short:"v" long:"verbose" description:"log verbose debug information"`
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *BlobstoreCreateFileCommand) Execute(args []string) error {
	if !cmd.Verbose {
		util.StopLogging()
	}

	name := "docker-proxy"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := createFile(name)
	if err != nil {
		return err
	}
	fmt.Println("Created file blobstore", name)
	return nil
}

type payload struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func createFile(name string) error {
	payload := payload{
		Path: name,
		Name: name,
	}

	return api.Post("v1/blobstores/file", payload, 204)
}
