package schemas


type TaggedAddresses struct {
	Lan     string `json:"lan"`
	LanIpv4 string `json:"lan_ipv4"`
	Wan     string `json:"wan"`
	WanIpv4 string `json:"wan_ipv4"`
}
type ServiceMeta struct {
	Broker    string `json:"broker"`
	Protocol  string `json:"protocol"`
	Registry  string `json:"registry"`
	Server    string `json:"server"`
	Transport string `json:"transport"`
}
type ServiceProxy             struct {
	Mode        string `json:"Mode"`
	MeshGateway struct {
	} `json:"MeshGateway"`
	Expose struct {
	} `json:"Expose"`
}
type LanIpv4 struct {
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}
type WanIpv4 struct {
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}
type ServiceTaggedAddresses struct {
	LanIpv4 LanIpv4 `json:"lan_ipv4"`
	WanIpv4 WanIpv4 `json:"wan_ipv4"`
}
type ServiceWeights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
}
type SchemaConsulCatalogService struct {
	ID              string `json:"ID"`
	Node            string `json:"Node"`
	Address         string `json:"Address"`
	Datacenter      string `json:"Datacenter"`
	TaggedAddresses TaggedAddresses `json:"TaggedAddresses"`
	NodeMeta struct {
		ConsulNetworkSegment string `json:"consul-network-segment"`
	} `json:"NodeMeta"`
	ServiceKind            string   `json:"ServiceKind"`
	ServiceID              string   `json:"ServiceID"`
	ServiceName            string   `json:"ServiceName"`
	ServiceTags            []string `json:"ServiceTags"`
	ServiceAddress         string   `json:"ServiceAddress"`
	ServiceTaggedAddresses ServiceTaggedAddresses `json:"ServiceTaggedAddresses"`
	ServiceWeights ServiceWeights `json:"ServiceWeights"`
	ServiceMeta ServiceMeta `json:"ServiceMeta"`
	ServicePort              int    `json:"ServicePort"`
	ServiceSocketPath        string `json:"ServiceSocketPath"`
	ServiceEnableTagOverride bool   `json:"ServiceEnableTagOverride"`
	ServiceProxy  ServiceProxy `json:"ServiceProxy"`
	ServiceConnect struct {
	} `json:"ServiceConnect"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}


