package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"org/sonatype/nx/blobstore"
	"org/sonatype/nx/cleanup"
	"org/sonatype/nx/db"
	"org/sonatype/nx/docker"
	"org/sonatype/nx/maven"
	"org/sonatype/nx/network"
	"org/sonatype/nx/raw"
	"org/sonatype/nx/security"
	"org/sonatype/nx/selector"
	"org/sonatype/nx/task"
	"org/sonatype/nx/util"
	"os"
)

func main() {
	var opts struct {
		Verbose                 func()                                  `short:"v" long:"verbose" description:"log verbose debug information"`
		BlobstoreCreateFile     blobstore.BlobstoreCreateFileCommand    `command:"blobstore-create-file"`
		CleanupCreatePolicy     cleanup.CleanupCreatePolicyCommand      `command:"cleanup-create-policy"`
		DbListAssets            db.DbListAssetsCommand                  `command:"db-list-assets"`
		DockerCreateHosted      docker.DockerCreateHostedCommand        `command:"docker-create-hosted"`
		DockerCreateProxy       docker.DockerCreateProxyCommand         `command:"docker-create-proxy"`
		MavenCreateHosted       maven.MavenCreateHostedCommand          `command:"maven-create-hosted"`
		MavenCreateProxy        maven.MavenCreateProxyCommand           `command:"maven-create-proxy"`
		NetworkWait             network.NetworkWaitCommand              `command:"network-wait"`
		RawHosted               raw.RawCreateHostedCommand              `command:"raw-create-hosted"`
		RawProxy                raw.RawCreateProxyCommand               `command:"raw-create-proxy"`
		SelectorCreate          selector.SelectorCreateCommand          `command:"selector-create"`
		SelectorCreatePrivilege selector.SelectorCreatePrivilegeCommand `command:"selector-create-privilege"`
		SecurityCreateRole      security.CreateRoleCommand              `command:"create-role"`
		SecurityCreateUser      security.CreateUserCommand              `command:"create-user"`
		SetAnonymous            security.SetAnonymousCommand            `command:"set-anonymous"`
		TaskList                task.TaskListCommand                    `command:"task-list"`
		TaskRun                 task.TaskRunCommand                     `command:"task-run"`
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
