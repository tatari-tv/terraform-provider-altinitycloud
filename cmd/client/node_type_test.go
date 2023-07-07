package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAltinityCloudClient_GetNodeType(t *testing.T) {
	// Create a mock server to handle the API request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the URL and method to ensure it matches the expected request
		if r.URL.Path != "/environment/test/nodetypes" {
			t.Errorf("Expected URL path: /environment/test/nodetypes, got: %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected request method: GET, got: %s", r.Method)
		}

		// Prepare a mock response
		response := NodeTypeData{
			NodeTypes: []NodeType{
				{
					Name:         "test-node",
					Scope:        "test-scope",
					Code:         "test-code",
					StorageClass: "test-storage",
					CPU:          "test-cpu",
					Memory:       "test-memory",
				},
			},
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			t.Fatalf("Failed to marshal response: %s", err)
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON)
	}))
	defer server.Close()

	// Create a AltinityCloudClient instance with the mock server URL
	client := &AltinityCloudClient{
		HTTPClient:  http.DefaultClient,
		APIEndpoint: server.URL,
	}

	// Call the method being tested
	nodeType, err := client.GetNodeType("test", "test-node")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// Define the expected NodeType
	expected := NodeType{
		Name:         "test-node",
		Scope:        "test-scope",
		Code:         "test-code",
		StorageClass: "test-storage",
		CPU:          "test-cpu",
		Memory:       "test-memory",
	}

	// Compare the returned NodeType with the expected value
	if !reflect.DeepEqual(nodeType, expected) {
		t.Errorf("Returned NodeType does not match expected value.\nExpected: %+v\nGot: %+v", expected, nodeType)
	}
}

func TestAltinityCloudClient_GetNodeTypes(t *testing.T) {
	// Create a mock server to handle the API request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the URL and method to ensure it matches the expected request
		if r.URL.Path != "/environment/test/nodetypes" {
			t.Errorf("Expected URL path: /environment/test/nodetypes, got: %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected request method: GET, got: %s", r.Method)
		}

		// Prepare a mock response
		response := NodeTypeData{
			NodeTypes: []NodeType{
				{
					Name:         "test-node1",
					Scope:        "test-scope1",
					Code:         "test-code1",
					StorageClass: "test-storage1",
					CPU:          "test-cpu1",
					Memory:       "test-memory1",
				},
				{
					Name:         "test-node2",
					Scope:        "test-scope2",
					Code:         "test-code2",
					StorageClass: "test-storage2",
					CPU:          "test-cpu2",
					Memory:       "test-memory2",
				},
			},
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			t.Fatalf("Failed to marshal response: %s", err)
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON)
	}))
	defer server.Close()

	// Create a AltinityCloudClient instance with the mock server URL
	client := &AltinityCloudClient{
		HTTPClient:  http.DefaultClient,
		APIEndpoint: server.URL,
	}

	// Call the method being tested
	nodeTypes, err := client.GetNodeTypes("test")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// Define the expected NodeTypeData
	expected := NodeTypeData{
		NodeTypes: []NodeType{
			{
				Name:         "test-node1",
				Scope:        "test-scope1",
				Code:         "test-code1",
				StorageClass: "test-storage1",
				CPU:          "test-cpu1",
				Memory:       "test-memory1",
			},
			{
				Name:         "test-node2",
				Scope:        "test-scope2",
				Code:         "test-code2",
				StorageClass: "test-storage2",
				CPU:          "test-cpu2",
				Memory:       "test-memory2",
			},
		},
	}

	// Compare the returned NodeTypeData with the expected value
	if !reflect.DeepEqual(nodeTypes, expected) {
		t.Errorf("Returned NodeTypeData does not match expected value.\nExpected: %+v\nGot: %+v", expected, nodeTypes)
	}
}

func TestAltinityCloudClient_CreateNodeType(t *testing.T) {
	// Create a mock server to handle the API request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the URL and method to ensure it matches the expected request
		if r.URL.Path != "/environment/test/nodetypes" {
			t.Errorf("Expected URL path: /environment/test/nodetypes, got: %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected request method: POST, got: %s", r.Method)
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %s", err)
		}

		// Verify the request parameters
		expectedParams := "name=test-node&scope=test-scope&code=test-code&storageClass=test-storage&cpu=test-cpu&memory=test-memory"
		if string(body) != expectedParams {
			t.Errorf("Expected request body: %s, got: %s", expectedParams, string(body))
		}

		// Prepare a mock response
		response := NodeTypeCreateResponse{
			Data: NodeType{
				Name:         "test-node",
				Scope:        "test-scope",
				Code:         "test-code",
				StorageClass: "test-storage",
				CPU:          "test-cpu",
				Memory:       "test-memory",
			},
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			t.Fatalf("Failed to marshal response: %s", err)
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON)
	}))
	defer server.Close()

	// Create a AltinityCloudClient instance with the mock server URL
	client := &AltinityCloudClient{
		HTTPClient:  http.DefaultClient,
		APIEndpoint: server.URL,
	}

	// Create a NodeType for testing
	nodeType := NodeType{
		Name:         "test-node",
		Scope:        "test-scope",
		Code:         "test-code",
		StorageClass: "test-storage",
		CPU:          "test-cpu",
		Memory:       "test-memory",
	}

	// Call the method being tested
	createdNodeType, err := client.CreateNodeType("test", nodeType)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// Define the expected NodeType
	expected := NodeType{
		Name:         "test-node",
		Scope:        "test-scope",
		Code:         "test-code",
		StorageClass: "test-storage",
		CPU:          "test-cpu",
		Memory:       "test-memory",
	}

	// Compare the created NodeType with the expected value
	if !reflect.DeepEqual(createdNodeType, expected) {
		t.Errorf("Created NodeType does not match expected value.\nExpected: %+v\nGot: %+v", expected, createdNodeType)
	}
}
