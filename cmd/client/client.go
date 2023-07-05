package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIEndpoint - Altinity.Cloud default API endpoint.
const APIEndpoint string = "https://acm.altinity.cloud/api"

// AltinityCloudClient is wrapper for http client and configs.
type AltinityCloudClient struct {
	HTTPClient  *http.Client
	APIEndpoint string
	APIToken    string
}

// NewClient - create new Altinity.Cloud client.
func NewClient(endpoint, token *string) (*AltinityCloudClient, error) {
	c := AltinityCloudClient{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Altinity.Cloud API endpoint
		APIEndpoint: APIEndpoint,
		APIToken:    "",
	}

	if endpoint != nil {
		c.APIEndpoint = *endpoint
	}

	if token != nil {
		c.APIToken = *token
	}

	return &c, nil
}

// doRequest - sends HTTP over the wire with correct headers and returns response.
func (c *AltinityCloudClient) doRequest(req *http.Request, authToken *string) ([]byte, error) { // nolint: unparam
	token := c.APIToken

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("X-Auth-Token", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("client: could not do request: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
