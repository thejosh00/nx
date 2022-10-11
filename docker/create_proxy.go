package docker

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
)

type DockerCreateProxyCommand struct {
	Pull          bool   `long:"pull" description:"enable pull replication"`
	Username      string `short:"u" long:"user" default:"admin" description:"username for authentication"`
	Password      string `short:"p" long:"password" default:"admin123" description:"password for authentication"`
	Remote        string `short:"r" long:"remote" default:"http://localhost:18001" description:"remote url of server to proxy"`
	RepositoryUrl string `long:"repositoryUrl" default:"http://localhost:8081/repository/docker-hosted" description:"Repository URL (if it's not the same as remote)"`
	Positional    struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *DockerCreateProxyCommand) Execute(args []string) error {
	name := "docker-proxy"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := createProxy(name, cmd)
	if err != nil {
		return err
	}
	fmt.Println("Created docker proxy", name)
	return nil
}

type docker struct {
	V1Enabled      bool `json:"v1Enabled"`
	ForceBasicAuth bool `json:"forceBasicAuth"`
	HttpPort       int  `json:"httpPort"`
}

type dockerProxy struct {
	IndexType string `json:"indexType"`
}

type payload struct {
	Name          string                   `json:"name"`
	Online        bool                     `json:"online"`
	Storage       *repomodel.Storage       `json:"storage,omitempty"`
	Proxy         *repomodel.Proxy         `json:"proxy,omitempty"`
	NegativeCache *repomodel.NegativeCache `json:"negativeCache,omitempty"`
	HttpClient    *repomodel.HttpClient    `json:"httpClient,omitempty"`
	Replication   *repomodel.Replication   `json:"replication,omitempty"`
	Docker        *docker                  `json:"docker,omitempty"`
	DockerProxy   *dockerProxy             `json:"dockerProxy,omitempty"`
}

func createProxy(name string, cmd *DockerCreateProxyCommand) error {
	payload := payload{
		Name:   name,
		Online: true,
		Storage: &repomodel.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: &repomodel.Proxy{
			RemoteUrl:      cmd.Remote,
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: &repomodel.NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		HttpClient: &repomodel.HttpClient{
			Blocked:   false,
			AutoBlock: true,
			Connection: &repomodel.Connection{
				UseTrustStore: false,
			},
		},
		Replication: &repomodel.Replication{
			PreemptivePullEnabled: cmd.Pull,
			RepositoryUrl:         cmd.RepositoryUrl,
		},
		Docker: &docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HttpPort:       18002,
		},
		DockerProxy: &dockerProxy{
			IndexType: "REGISTRY",
		},
	}

	if cmd.Username != "" || cmd.Password != "" {
		payload.HttpClient.Authentication = &repomodel.Authentication{
			Type:     "username",
			Username: cmd.Username,
			Password: cmd.Password,
		}
	}

	return api.Post("v1/repositories/docker/proxy", payload, 201)
}
