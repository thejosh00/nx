package maven

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
)

type MavenCreateProxyCommand struct {
	Pull       bool   `long:"pull" description:"enable pull replication"`
	Username   string `short:"u" long:"user" default:"admin" description:"username for authentication"`
	Password   string `short:"p" long:"password" default:"admin123" description:"password for authentication"`
	Remote     string `short:"r" long:"remote" default:"http://localhost:8081/repository/maven-hosted" description:"remote url of server to proxy"`
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *MavenCreateProxyCommand) Execute(args []string) error {
	name := "maven-proxy"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	err := createProxy(name, cmd)
	if err != nil {
		return err
	}
	fmt.Println("Created maven proxy repository", name)
	return nil
}

type proxyPayload struct {
	Name          string                   `json:"name"`
	Online        bool                     `json:"online"`
	Proxy         *repomodel.Proxy         `json:"proxy,omitempty"`
	NegativeCache *repomodel.NegativeCache `json:"negativeCache,omitempty"`
	HttpClient    *repomodel.HttpClient    `json:"httpClient,omitempty"`
	Storage       *repomodel.Storage       `json:"storage,omitempty"`
	Replication   *repomodel.Replication   `json:"replication,omitempty"`
	Maven         maven                    `json:"maven"`
}

func createProxy(name string, cmd *MavenCreateProxyCommand) error {
	payload := proxyPayload{
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
		},
		Maven: maven{
			VersionPolicy:      "MIXED",
			LayoutPolicy:       "STRICT",
			ContentDisposition: "ATTACHMENT",
		},
	}

	if cmd.Username != "" || cmd.Password != "" {
		payload.HttpClient.Authentication = &repomodel.Authentication{
			Type:     "username",
			Username: cmd.Username,
			Password: cmd.Password,
		}
	}

	return api.Post("v1/repositories/maven/proxy", payload, 201)
}
