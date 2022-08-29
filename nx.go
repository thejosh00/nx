package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"org/sonatype/nx/blobstore"
	"org/sonatype/nx/cleanup"
	"org/sonatype/nx/docker"
	"org/sonatype/nx/maven"
	"org/sonatype/nx/network"
	"org/sonatype/nx/raw"
	"org/sonatype/nx/security"
	"org/sonatype/nx/util"
	"os"
)

func main() {
	var opts struct {
		Verbose             func()                               `short:"v" long:"verbose" description:"log verbose debug information"`
		BlobstoreCreateFile blobstore.BlobstoreCreateFileCommand `command:"blobstore-create-file"`
		CleanupCreatePolicy cleanup.CleanupCreatePolicyCommand   `command:"cleanup-create-policy"`
		DockerCreateProxy   docker.DockerCreateProxyCommand      `command:"docker-create-proxy"`
		MavenCreateHosted   maven.MavenCreateHostedCommand       `command:"maven-create-hosted"`
		MavenCreateProxy    maven.MavenCreateProxyCommand        `command:"maven-create-proxy"`
		NetworkWait         network.NetworkWaitCommand           `command:"network-wait"`
		RawHosted           raw.RawCreateHostedCommand           `command:"raw-create-hosted"`
		RawProxy            raw.RawCreateProxyCommand            `command:"raw-create-proxy"`
		SetAnonymous        security.SetAnonymousCommand         `command:"set-anonymous"`
	}

	util.StopLogging()

	opts.Verbose = func() {
		util.StartLogging()
	}

	var parser = flags.NewParser(&opts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
