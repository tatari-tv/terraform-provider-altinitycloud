package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetNodeTypes - Returns list of node types from Altinity.Cloud API.
func (c *AltinityCloudClient) GetNodeTypes(envID string) (NodeTypeData, error) {
	requestURL := fmt.Sprintf("%s/environment/%s/nodetypes", c.APIEndpoint, envID)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return NodeTypeData{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Printf("client: could not make request: %s\n", err)
		return NodeTypeData{}, err
	}

	nt := NodeTypeData{}
	err = json.Unmarshal(body, &nt)
	if err != nil {
		fmt.Printf("client: could not unmarshal data: %s\n", err)
		return NodeTypeData{}, err
	}

	return nt, nil
}

// GetNodeType - Returns node type by name from Altinity.Cloud API.
func (c *AltinityCloudClient) GetNodeType(envID, name string) (NodeType, error) {
	nts, err := c.GetNodeTypes(envID)
	if err != nil {
		fmt.Printf("client: could not get node types: %s\n", err)
		return NodeType{}, fmt.Errorf("client: could not get node type: %s", err)
	}

	// find the node type by name
	for _, nt := range nts.NodeTypes {
		if nt.Name == name {
			return nt, nil
		}
	}

	// return empty node type if not found
	return NodeType{}, nil
}

func (c *AltinityCloudClient) CreateNodeType(envID string, nt NodeType) (NodeType, error) {
	// build the POST request
	requestURL := fmt.Sprintf("%s/environment/%s/nodetypes", c.APIEndpoint, envID)
	req, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return NodeType{}, err
	}

	// add the query params
	q := req.URL.Query()
	q.Add("name", nt.Name)
	q.Add("scope", nt.Scope)
	q.Add("code", nt.Code)
	q.Add("storageClass", nt.StorageClass)
	q.Add("cpu", nt.CPU)
	q.Add("memory", nt.Memory)

	// add optional params pool if not null or empty string
	if len(nt.Pool) > 0 {
		q.Add("pool", nt.Pool)
	}

	// add optional params nodeSelector if not null or empty string
	if len(nt.NodeSelector) > 0 {
		q.Add("nodeSelector", nt.NodeSelector)
	}

	// add optional params tolerations if not null or empty string
	if len(nt.Tolerations) > 0 {
		// marshall the tolerations to a string
		tolerations, err := json.Marshal(nt.Tolerations)
		if err != nil {
			fmt.Printf("client: could not marshal tolerations: %s\n", err)
			return NodeType{}, err
		}
		q.Add("tolerations", string(tolerations))
	}

	req.URL.RawQuery = q.Encode()

	// make the request
	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Printf("client: could not make request: %s\n", err)
		return NodeType{}, err
	}

	// unmarshal the response
	ntcr := NodeTypeCreateResponse{}
	err = json.Unmarshal(body, &ntcr)
	if err != nil {
		fmt.Printf("client: could not unmarshal data: %s\n", err)
		return NodeType{}, err
	}

	return ntcr.Data, nil
}
