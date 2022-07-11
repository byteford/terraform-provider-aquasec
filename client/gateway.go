package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// User represents a local Aqua user
type Gateway struct {
	ID           string `json:"id"`
	Logical_Name string `json:"logicalname"`
	Description  string `json:"description"`
	Version      string `json:"version"`
	Host_Name    string `json:"hostname"`
	SSH_Address  string `json:"public_address"`
	GRPC_Address string `json:"grpc_address"`
	Status       string `json:"status"`
}

// GetUser - returns single Aqua gateway
func (cli *Client) GetGateway(name string) (*Gateway, error) {
	var err error
	var response Gateway
	cli.gorequest.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/servers/%s", name)
	events, body, errs := cli.gorequest.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
		err = fmt.Errorf("error calling %s", apiPath)
		return nil, err
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetGateway from %s%s, %v ", cli.url, apiPath, err)
			return nil, err
		}
	}
	if response.ID == "" {
		err = fmt.Errorf("gateway not found: %s", name)
		return nil, err
	}
	return &response, err
}

// GetUsers - returns all Aqua gateways
func (cli *Client) GetGateways() ([]Gateway, error) {
	var err error
	var response []Gateway
	request := cli.gorequest
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v1/servers")
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		err = fmt.Errorf("error calling %s", apiPath)
		return nil, err
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetGateways from %s%s, %v ", cli.url, apiPath, err)
			return nil, errors.Wrap(err, "could not unmarshal gateways response")
		}
	}
	return response, err
}
