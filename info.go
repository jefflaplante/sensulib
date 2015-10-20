/*
List the Sensu version and the transport and Redis connection information. This is the same information that /health uses to determine system health.
https://sensuapp.org/docs/0.20/api-info
*/

package sensu

import (
	"net/http"
)

// InfoURI for the sensu API
const InfoURI string = "/info"

// GetInfo queries service info
// The Sensu API responds with a nested JSON object. Unmarshal it into a map.
func (c *API) GetInfo(out interface{}) (*http.Response, error) {
	resp, err := c.get(InfoURI, out)
	return resp, err
}
