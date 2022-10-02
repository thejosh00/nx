package repomodel

type Storage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
}

type HostedStorage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
}

type Proxy struct {
	RemoteUrl      string `json:"remoteUrl"`
	RepositoryUrl  string `json:"repositoryUrl"`
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge int    `json:"metadataMaxAge"`
}

type NegativeCache struct {
	Enabled    bool `json:"enabled"`
	TimeToLive int  `json:"timeToLive"`
}

type Connection struct {
	UseTrustStore bool `json:"useTrustStore"`
}

type Authentication struct {
	Type       string `json:"type"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Preemptive bool   `json:"preemptive"`
}

type HttpClient struct {
	Blocked        bool           `json:"blocked"`
	AutoBlock      bool           `json:"autoBlock"`
	Connection     Connection     `json:"connection"`
	Authentication Authentication `json:"authentication"`
}

type Replication struct {
	PreemptivePullEnabled bool `json:"preemptivePullEnabled"`
}
