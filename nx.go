package main

import (
	"github.com/jessevdk/go-flags"
	"org/sonatype/nx/blobstore"
	"org/sonatype/nx/docker"
	"org/sonatype/nx/network"
	"org/sonatype/nx/raw"
	"org/sonatype/nx/security"
	"os"
)

func main() {
	var opts struct {
		BlobstoreCreateFile blobstore.BlobstoreCreateFileCommand `command:"blobstore-create-file"`
		DockerCreateProxy   docker.DockerCreateProxyCommand      `command:"docker-create-proxy"`
		NetworkWait         network.NetworkWaitCommand           `command:"network-wait"`
		RawHosted           raw.RawCreateHostedCommand           `command:"raw-create-hosted"`
		RawProxy            raw.RawCreateProxyCommand            `command:"raw-create-proxy"`
		SetAnonymous        security.SetAnonymousCommand         `command:"set-anonymous"`
	}

	var parser = flags.NewParser(&opts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		os.Exit(-1)
	}
}
