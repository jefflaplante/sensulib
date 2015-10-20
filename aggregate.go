/*
List and delete check aggregates.
This endpoint provides the information needed to monitor a collection of machines running a service.
https://sensuapp.org/docs/0.20/api-aggregates
*/

package sensu

import (
	"net/http"
	"strconv"
	"strings"
)

// AggregatesURI for the sensu API
const AggregatesURI string = "/aggregates"

// Aggregate represents an Aggregate object in the Sensu API
type Aggregate struct {
	Ok       int `json:"ok"`
	Warning  int `json:"warning"`
	Critical int `json:"critical"`
	Unknown  int `json:"unkown"`
	Total    int `json:"total"`
}

// GetAggregates returns all aggregates
func (c *API) GetAggregates(out interface{}) (*http.Response, error) {
	resp, err := c.get(AggregatesURI, out)
	return resp, err
}

// GetAggregatesByCheck returns all aggregates for the given check name
func (c *API) GetAggregatesByCheck(out interface{}, checkName string) (*http.Response, error) {
	s := []string{AggregatesURI, checkName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// RemoveAggregates removes all aggregates by the given check name
func (c *API) RemoveAggregates(checkName string) (*http.Response, error) {
	s := []string{AggregatesURI, checkName}
	uri := strings.Join(s, "/")
	resp, err := c.delete(uri)
	return resp, err
}

// GetAggregateByCheckIssued returns the aggregate for the given check name and issued ID
func (c *API) GetAggregateByCheckIssued(out interface{}, checkName string, issued int) (*http.Response, error) {
	s := []string{AggregatesURI, checkName, strconv.Itoa(issued)}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}
