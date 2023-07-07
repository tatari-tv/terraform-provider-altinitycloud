package client

// NodeTypeData - list of NodeType types.
type NodeTypeData struct {
	NodeTypes []NodeType `json:"data"`
}

// NodeType - node type model.
type NodeType struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Scope        string       `json:"scope"`
	Code         string       `json:"code"`
	Pool         string       `json:"pool"`
	StorageClass string       `json:"storageClass"`
	CPU          string       `json:"cpu"`
	Memory       string       `json:"memory"`
	ExtraSpec    string       `json:"extraSpec,omitempty"`
	Tolerations  []Toleration `json:"tolerations,omitempty"`
	NodeSelector string       `json:"nodeSelector,omitempty"`
	CPUAlloc     string       `json:"cpu_alloc,omitempty"`
	MemoryAlloc  string       `json:"memory_alloc,omitempty"`
}

// Toleration - node type toleration model.
type Toleration struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Effect   string `json:"effect"`
}

// NodeTypeCreateResponse - response from create node type.
type NodeTypeCreateResponse struct {
	Metadata struct {
		Changed bool `json:"changed"`
	} `json:"metadata"`
	Data NodeType `json:"data"`
}
