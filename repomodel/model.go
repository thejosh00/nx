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

type HttpClient struct {
	Blocked    bool       `json:"blocked"`
	AutoBlock  bool       `json:"autoBlock"`
	Connection Connection `json:"connection"`
}

type Replication struct {
	PreemptivePullEnabled bool `json:"preemptivePullEnabled"`
}
