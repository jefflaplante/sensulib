/*
List locally defined checks and request executions.
https://sensuapp.org/docs/0.20/api-checks
*/

package sensu

import (
	"net/http"
	"strings"
)

// ChecksURI constant
const ChecksURI string = "/checks"

// RequestURI constant
const RequestURI string = "/request"

// Check represents a Check in the Sensu API
type Check struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Subscribers []string `json:"subscribers"`
	Interval    int      `json:"interval"`
	Issued      int      `json:"issued"`
	Executed    int      `json:"executed"`
	Output      string   `json:"output"`
	Status      int      `json:"status"`
	Duration    int      `json:"duration"`
}

// Request represents a Request in the Sensu API
type Request struct {
	CheckName   string   `json:"check"`
	Subscribers []string `json:"subscribers"`
}

// GetChecks gets all checks
func (c *API) GetChecks(out interface{}) (*http.Response, error) {
	resp, err := c.get(ChecksURI, out)
	return resp, err
}

// GetCheck gets an individual check by name
func (c *API) GetCheck(out interface{}, checkName string) (*http.Response, error) {
	s := []string{ChecksURI, checkName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// CheckRequest requests a check be performed, by name as expressed in the payload struct
func (c *API) CheckRequest(payload interface{}) (*http.Response, error) {
	resp, err := c.post(RequestURI, payload)
	return resp, err
}
