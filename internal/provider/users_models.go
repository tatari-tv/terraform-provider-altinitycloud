package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

// UsersDataSourceModel describes the Users model for data sources.
type UsersDataSourceModel struct {
	ClusterID types.String `tfsdk:"cluster_id"`
	Users     []UserModel  `tfsdk:"users"`
}

// UserModel describes a user in an Altintiy.Cloud cluster.
type UserModel struct {
	ID               types.String `tfsdk:"id"`
	Login            types.String `tfsdk:"login"`
	Password         types.String `tfsdk:"password"`
	Networks         types.String `tfsdk:"networks"`
	Databases        types.String `tfsdk:"databases"`
	ProfileID        types.String `tfsdk:"profile_id"`
	QuotaID          types.String `tfsdk:"quota_id"`
	AccessManagement types.Bool   `tfsdk:"access_management"`
	System           types.Bool   `tfsdk:"system"`
}
