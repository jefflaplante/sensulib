/*
Check the status of the API’s transport & Redis connections, and query the transport’s status. (consumer and message counts)
https://sensuapp.org/docs/0.20/api-health
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
