package raw

import (
	"fmt"
	"log"
	"org/sonatype/nx/api"
	"org/sonatype/nx/repomodel"
	"org/sonatype/nx/util"
)

type RawCreateProxyCommand struct {
	Verbose bool `short:"v" long:"verbose" description:"log verbose debug information"`
}

func (cmd *RawCreateProxyCommand) Execute(args []string) error {
	if !cmd.Verbose {
		log.SetOutput(new(util.NoLogger))
	}

	err := createProxy("raw-proxy")
	if err != nil {
		return err
	}
	fmt.Println("Created raw proxy repository")
	return nil
}

type proxyPayload struct {
	Name          string                  `json:"name"`
	Online        bool                    `json:"online"`
	Proxy         repomodel.Proxy         `json:"proxy"`
	NegativeCache repomodel.NegativeCache `json:"negativeCache"`
	HttpClient    repomodel.HttpClient    `json:"httpClient"`
	Storage       repomodel.Storage       `json:"storage"`
}

func createProxy(name string) error {
	payload := proxyPayload{
		Name:   name,
		Online: true,
		Storage: repomodel.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Proxy: repomodel.Proxy{
			RemoteUrl:      "http://localhost:8081/repository/raw-hosted",
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
	}

	return api.Post("v1/repositories/raw/proxy", payload, 201)
}
