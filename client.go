/*
List and delete client(s) information.
https://sensuapp.org/docs/0.20/api-clients
*/

package sensu

import (
	"net/http"
	"strings"
)

// ClientsURI Constant
const clientsURI string = "/clients"

// Client represents a Client object from the Sensu API
type Client struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Version       string   `json:"version"`
	Timestamp     int      `json:"timestamp"`
	Environment   string   `json:"environment"`
}

// ClientHistory represents a ClientHistory object from the Sensu API
type ClientHistory struct {
	Check         string                 `json:"check"`
	History       []int                  `json:"history"`
	LastExecution int                    `json:"last_execution"`
	LastStatus    int                    `json:"last_status"`
	LastResult    map[string]interface{} `json:"last_result"`
}

// GetClients gets all clients
// sets 'out' as an array of Client structs as JSON
func (c *API) GetClients(out interface{}) (*http.Response, error) {
	resp, err := c.get(clientsURI, out)
	return resp, err
}

// GetClient gets an individual client
// sets 'out' as a Client struct as JSON
func (c *API) GetClient(out interface{}, clientName string) (*http.Response, error) {
	s := []string{clientsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// GetClientHistory gets a client's history
func (c *API) GetClientHistory(out interface{}, clientName string) (*http.Response, error) {
	s := []string{clientsURI, clientName, "history"}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// RemoveClient removes a client by name
func (c *API) RemoveClient(clientName string) (*http.Response, error) {
	s := []string{clientsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.delete(uri)
	return resp, err
}

// CreateClient creates a client
// POST a Client object to the API
func (c *API) CreateClient(payload interface{}) (*http.Response, error) {
	resp, err := c.post(clientsURI, payload)
	return resp, err
}
