package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// NodeTypeDataSourceModel describes the NodeTypes source NodeTypes model for data sources.
// @note: these attributes are at the same level because of this issue: https://github.com/hashicorp/terraform-plugin-framework/issues/191
type NodeTypeDataSourceModel struct {
	EnvID        types.String      `tfsdk:"env_id"`
	ID           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	Scope        types.String      `tfsdk:"scope"`
	Code         types.String      `tfsdk:"code"`
	Pool         types.String      `tfsdk:"pool"`
	StorageClass types.String      `tfsdk:"storage_class"`
	CPU          types.String      `tfsdk:"cpu"`
	Memory       types.String      `tfsdk:"memory"`
	ExtraSpec    types.String      `tfsdk:"extra_spec"`
	Tolerations  []TolerationModel `tfsdk:"tolerations"`
	NodeSelector types.String      `tfsdk:"node_selector"`
	CPUAlloc     types.String      `tfsdk:"cpu_alloc"`
	MemoryAlloc  types.String      `tfsdk:"memory_alloc"`
}

// NodeTypeResourceModel - describes the NodeTypes source NodeTypes model for resources.
type NodeTypeResourceModel struct {
	EnvID       types.String  `tfsdk:"env_id"`
	NodeType    NodeTypeModel `tfsdk:"node_type"`
	LastUpdated types.String  `tfsdk:"last_updated"`
}

// NodeTypeModel - node type datasource representation.
type NodeTypeModel struct {
	ID           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	Scope        types.String      `tfsdk:"scope"`
	Code         types.String      `tfsdk:"code"`
	Pool         types.String      `tfsdk:"pool"`
	StorageClass types.String      `tfsdk:"storage_class"`
	CPU          types.String      `tfsdk:"cpu"`
	Memory       types.String      `tfsdk:"memory"`
	ExtraSpec    types.String      `tfsdk:"extra_spec"`
	NodeSelector types.String      `tfsdk:"node_selector"`
	Tolerations  []TolerationModel `tfsdk:"tolerations"`
	CPUAlloc     types.String      `tfsdk:"cpu_alloc"`
	MemoryAlloc  types.String      `tfsdk:"memory_alloc"`
}

// TolerationModel - Kubernetes tolerations.
type TolerationModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Effect   types.String `tfsdk:"effect"`
	Value    types.String `tfsdk:"value"`
}
