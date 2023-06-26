package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	endpoint := "https://example.org/api"
	token := "notsosecret"

	valid, err := NewClient(&endpoint, &token)
	if err != nil {
		t.Fatalf(`NewClient(&endpoint, &token), want nil got %v`, err)
	}

	assert.Equal(t, endpoint, valid.APIEndpoint, "Altiniy.Cloud endpoints should match")
	assert.Equal(t, token, valid.APIToken, "Altiniy.Cloud tokens string should match")
}

func TestEmptyClient(t *testing.T) {
	valid, err := NewClient(nil, nil)
	if err != nil {
		t.Fatalf(`NewClient(&endpoint, &token), want nil got %v`, err)
	}

	assert.Equal(t, "https://acm.altinity.cloud/api", valid.APIEndpoint, "Altiniy.Cloud endpoints should match")
	assert.Equal(t, "", valid.APIToken, "Altiniy.Cloud tokens string should match")
}
