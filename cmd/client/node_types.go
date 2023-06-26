package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetNodeTypes - Returns list of node types from Altinity.Cloud API.
func (c *AltinityCloudClient) GetNodeTypes(envID string) (NodeTypeData, error) {
	requestURL := fmt.Sprintf("%s/environment/%s/nodetypes", c.APIEndpoint, "652")
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
