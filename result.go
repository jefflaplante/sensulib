/*
List current check results.
https://sensuapp.org/docs/0.20/api-results
*/

package sensu

import (
	"fmt"
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

// CheckResult struct contains the "bare minimum" needed to POST to sensu-api's /results route.
type CheckResult struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Output  string `json:"output"`
	Status  int    `json:"status"`
	Handler string `json:"handler"`
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

// PostCheckResult  POST's check results to sensu-api's /results route.
//  see https://sensuapp.org/docs/latest/api-results#results-post
func (c *API) PostCheckResult(result CheckResult) error {
	res, err := c.post(ResultsURI, result)
	if res.StatusCode != 202 {
		return fmt.Errorf("Could not post result to Sensu-API. Expected 202 but received %d, %s", res.StatusCode, res.Status)
	}

	return err
}
