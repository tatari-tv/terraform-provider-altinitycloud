package client

// NodeTypeData - list of NodeType types.
type NodeTypeData struct {
	Data []NodeType `json:"data"`
}

// NodeType - node type model.
type NodeType struct {
	ID            string       `json:"id"`
	Scope         string       `json:"scope"`
	Code          string       `json:"code"`
	Name          string       `json:"name"`
	Pool          string       `json:"pool"`
	StorageClass  string       `json:"storageClass"`
	CPU           string       `json:"cpu"`
	Memory        string       `json:"memory"`
	IDEnvironment string       `json:"id_environment"`
	ExtraSpec     string       `json:"extraSpec"`
	Tolerations   []Toleration `json:"tolerations"`
	NodeSelector  string       `json:"nodeSelector"`
	CPUAlloc      string       `json:"cpu_alloc"`
	MemoryAlloc   string       `json:"memory_alloc"`
}

// Toleration - node type toleration model.
type Toleration struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
	Effect   string `json:"effect"`
}
