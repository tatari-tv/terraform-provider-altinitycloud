package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *AltinityCloudClient) GetUsers(clusterID string) (UserData, error) {
	// build the GET request
	requestURL := fmt.Sprintf("%s/cluster/%s/users", c.APIEndpoint, clusterID)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return UserData{}, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Printf("client: could not make request: %s\n", err)
		return UserData{}, err
	}

	ud := UserData{}
	err = json.Unmarshal(body, &ud)
	if err != nil {
		fmt.Printf("client: could not unmarshal data: %s\n", err)
		return UserData{}, err
	}

	return ud, nil
}
