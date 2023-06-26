package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// NodeTypeDataSourceModel describes the Data source Data model.
type NodeTypeDataSourceModel struct {
	EnvId types.String    `tfsdk:"env_id"`
	Id    types.String    `tfsdk:"id"`
	Data  []NodeTypeModel `tfsdk:"data"`
}

type NodeTypeModel struct {
	ID            types.Int64               `tfsdk:"id"`
	Scope         types.String              `tfsdk:"scope"`
	Code          types.String              `tfsdk:"code"`
	Name          types.String              `tfsdk:"name"`
	Pool          types.String              `tfsdk:"pool"`
	StorageClass  types.String              `tfsdk:"storage_class"`
	CPU           types.String              `tfsdk:"cpu"`
	Memory        types.String              `tfsdk:"memory"`
	IDEnvironment types.String              `tfsdk:"id_environment"`
	ExtraSpec     types.String              `tfsdk:"extra_spec"`
	Tolerations   []NodeTypeTolerationModel `tfsdk:"tolerations"`
	NodeSelector  types.String              `tfsdk:"node_selector"`
	CPUAlloc      types.String              `tfsdk:"cpu_alloc"`
	MemoryAlloc   types.String              `tfsdk:"memory_alloc"`
}

// NodeTypeTolerationModel - Kubernetes tolerations.
type NodeTypeTolerationModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Effect   types.String `tfsdk:"effect"`
	Value    types.String `tfsdk:"value"`
}
