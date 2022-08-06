package docker

import (
	"fmt"
	"log"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
	"org/sonatype/nx/util"
)

type DockerCreateProxyCommand struct {
	Verbose bool `short:"v" long:"verbose" description:"log verbose debug information"`
}

func (cmd *DockerCreateProxyCommand) Execute(args []string) error {
	if !cmd.Verbose {
		log.SetOutput(new(util.NoLogger))
	}

	err := createProxy("docker-proxy", 18001)
	if err != nil {
		return err
	}
	fmt.Println("Created docker proxy")
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
	Name          string                  `json:"name"`
	Online        bool                    `json:"online"`
	Storage       repomodel.Storage       `json:"storage"`
	Proxy         repomodel.Proxy         `json:"proxy"`
	NegativeCache repomodel.NegativeCache `json:"negativeCache"`
	HttpClient    repomodel.HttpClient    `json:"httpClient"`
	Docker        docker                  `json:"docker"`
	DockerProxy   dockerProxy             `json:"dockerProxy"`
}

func createProxy(name string, port int) error {
	payload := payload{
		Name:   name,
		Online: true,
		Storage: repomodel.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: repomodel.Proxy{
			RemoteUrl:      "https://registry-1.docker.io",
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
		},
		NegativeCache: repomodel.NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		HttpClient: repomodel.HttpClient{
			Blocked:   false,
			AutoBlock: true,
			Connection: repomodel.Connection{
				UseTrustStore: false,
			},
		},
		Docker: docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HttpPort:       port,
		},
		DockerProxy: dockerProxy{
			IndexType: "HUB",
		},
	}

	return api.Post("v1/repositories/docker/proxy", payload, 201)
}
