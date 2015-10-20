/*
Check the status of the API’s transport & Redis connections, and query the transport’s status. (consumer and message counts)
https://sensuapp.org/docs/0.20/api-health
*/

package sensu

import ()

// HealthURI for the sensu API
const HealthURI string = "/health"

// GetHealth queries service health
// Endpoint Returns 204 Success (healthy service)
// Endpoint Returns 503 Error (unhealthy service)
// Function returns a boolean (true/false) based on health
func (c *API) GetHealth() (bool, error) {
	uri := HealthURI
	resp, err := c.get(uri, nil)

	if resp.StatusCode == 204 {
		return true, err
	}
	return false, err
}
