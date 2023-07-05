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

func (c *AltinityCloudClient) UpdateNodeType(envID string, nodeType NodeType) (NodeType, error) {
	// get the node type
	n, err := c.GetNodeType(envID, nodeType.Name)
	if err != nil {
		fmt.Printf("client: could not get node type: %s\n", err)
		return NodeType{}, err
	}

	// build the POST request
	requestURL := fmt.Sprintf("%s/nodetype/%s", c.APIEndpoint, n.ID)
	req, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return NodeType{}, err
	}

	// add the query params
	q := req.URL.Query()
	q.Add("name", nodeType.Name)
	q.Add("scope", nodeType.Scope)
	q.Add("code", nodeType.Code)
	q.Add("storageClass", nodeType.StorageClass)
	q.Add("cpu", nodeType.CPU)
	q.Add("memory", nodeType.Memory)

	// add optional params pool if not null or empty string
	if len(nodeType.Pool) > 0 {
		q.Add("pool", nodeType.Pool)
	}

	// add optional params nodeSelector if not null or empty string
	if len(nodeType.NodeSelector) > 0 {
		q.Add("nodeSelector", nodeType.NodeSelector)
	}

	// add optional params tolerations if not null or empty string
	if len(nodeType.Tolerations) > 0 {
		// marshall the tolerations to a string
		tolerations, err := json.Marshal(nodeType.Tolerations)
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

func (c *AltinityCloudClient) DeleteNodeType(ID string) error {
	// build the POST request
	requestURL := fmt.Sprintf("%s/nodetype/%s", c.APIEndpoint, ID)
	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	// make the request
	_, err = c.doRequest(req, nil)
	if err != nil {
		fmt.Printf("client: could not make request: %s\n", err)
		return err
	}

	return nil
}
