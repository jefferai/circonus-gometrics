package api

// abstracted in preparation of separate circonus-api-go package

import (
	"encoding/json"
	"fmt"
)

type BrokerDetail struct {
	CN      string   `json:"cn"`
	IP      string   `json:"ipaddress"`
	MinVer  int      `json:"minimum_version_required"`
	Modules []string `json:"modules"`
	Port    int      `json:"port"`
	Skew    string   `json:"skew"`
	Status  string   `json:"status"`
	Version int      `json:"version"`
}

type Broker struct {
	Cid       string         `json:"_cid"`
	Details   []BrokerDetail `json:"_details"`
	Latitude  string         `json:"_latitude"`
	Longitude string         `json:"_longitude"`
	Name      string         `json:"_name"`
	Tags      []string       `json:"_tags"`
	Type      string         `json:"_type"`
}

// Use Circonus API to retrieve  a specific broker by Broker Group ID
func (a *Api) FetchBrokerById(id int) (*Broker, error) {
	cid := fmt.Sprintf("/broker/%d", id)
	return a.FetchBrokerByCid(cid)
}

// Use Circonus API to retreive a broker by CID
func (a *Api) FetchBrokerByCid(cid string) (*Broker, error) {
	result, err := a.Get(cid)
	if err != nil {
		return nil, err
	}

	response := new(Broker)
	if err := json.Unmarshal(result, &response); err != nil {
		return nil, err
	}

	return response, nil

}

// Use Circonus API to retreive a list of brokers which have a specific tag
func (a *Api) FetchBrokerListByTag(searchTag string) ([]Broker, error) {
	result, err := a.Get(fmt.Sprintf("/broker?f__tags_has=%s", searchTag))
	if err != nil {
		return nil, err
	}

	var response []Broker
	json.Unmarshal(result, &response)

	return response, nil
}

// Use Circonus API to retreive a list of brokers
func (a *Api) FetchBrokerList() ([]Broker, error) {
	result, err := a.Get("/broker")
	if err != nil {
		return nil, err
	}

	var response []Broker
	json.Unmarshal(result, &response)

	return response, nil
}
