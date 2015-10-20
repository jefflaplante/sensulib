/*
List current check results.
https://sensuapp.org/docs/0.20/api-results
*/

package sensu

import (
	"net/http"
	"strings"
)

// ResultsURI for the sensu API
const ResultsURI string = "/results"

// Result represents a Result object in the Sensu API
type Result struct {
	ClientName string `json:"client"`
	Check      Check  `json:"check"`
}

// GetResults gets all results
func (c *API) GetResults(out interface{}) (*http.Response, error) {
	resp, err := c.get(ResultsURI, out)
	return resp, err
}

// GetResultsByClient gets all results for the named client
func (c *API) GetResultsByClient(out interface{}, clientName string) (*http.Response, error) {
	s := []string{ResultsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// GetResultByClientCheck gets the result object for the named client and check
func (c *API) GetResultByClientCheck(out interface{}, clientName string, checkName string) (*http.Response, error) {
	s := []string{ResultsURI, clientName, checkName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}
