package provider

import (
	"github.com/stretchr/testify/assert"
	"github.com/tatari-tv/terraform-provider-altinitycloud/cmd/client"
	"testing"
)

func TestMapNodeTypeToNodeTypeResource(t *testing.T) {
	nt1 := client.NodeType{
		ID:           "123",
		Name:         "test",
		Scope:        "ClickHouse",
		Code:         "test",
		StorageClass: "gp3",
		CPU:          "1",
		Memory:       "1",
		Pool:         "test",
		Tolerations:  nil,
	}
	nt1.Tolerations = append(nt1.Tolerations, client.Toleration{
		Key:      "test",
		Operator: "test",
		Value:    "test",
		Effect:   "test",
	})

	ntr1 := mapNodeTypeToNodeTypeResponse(nt1)
	assert.Equal(t, ntr1.Name.ValueString(), nt1.Name)
	assert.Equal(t, ntr1.Scope.ValueString(), nt1.Scope)
	assert.Equal(t, ntr1.Code.ValueString(), nt1.Code)
	assert.Equal(t, ntr1.StorageClass.ValueString(), nt1.StorageClass)
	assert.Equal(t, ntr1.CPU.ValueString(), nt1.CPU)
	assert.Equal(t, ntr1.Memory.ValueString(), nt1.Memory)
	assert.Equal(t, ntr1.Pool.ValueString(), nt1.Pool)
	assert.Equal(t, ntr1.Tolerations[0].Key.ValueString(), nt1.Tolerations[0].Key)
	assert.Equal(t, ntr1.Tolerations[0].Operator.ValueString(), nt1.Tolerations[0].Operator)
	assert.Equal(t, ntr1.Tolerations[0].Value.ValueString(), nt1.Tolerations[0].Value)
	assert.Equal(t, ntr1.Tolerations[0].Effect.ValueString(), nt1.Tolerations[0].Effect)

	nt2 := client.NodeType{
		ID:           "123",
		Name:         "test",
		Scope:        "ClickHouse",
		Code:         "test",
		StorageClass: "gp3",
		CPU:          "1",
		Memory:       "1",
		Pool:         "test",
	}

	nrt2 := mapNodeTypeToNodeTypeResponse(nt2)
	assert.Equal(t, nrt2.Name.ValueString(), nt1.Name)
	assert.Equal(t, nrt2.Scope.ValueString(), nt1.Scope)
	assert.Equal(t, nrt2.Code.ValueString(), nt1.Code)
	assert.Equal(t, nrt2.StorageClass.ValueString(), nt1.StorageClass)
	assert.Equal(t, nrt2.CPU.ValueString(), nt1.CPU)
	assert.Equal(t, nrt2.Memory.ValueString(), nt1.Memory)
	assert.Equal(t, nrt2.Pool.ValueString(), nt1.Pool)
	assert.Equal(t, nrt2.NodeSelector.ValueString(), "")
	assert.Equal(t, len(nrt2.Tolerations), 0)
}
