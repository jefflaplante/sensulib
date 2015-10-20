/*
Create, list, and delete stashes (JSON documents). The stashes endpoint is an HTTP key/value data store.
https://sensuapp.org/docs/0.20/api-stashes
*/

package sensu

import (
	"net/http"
	"strings"
)

// StashesURI for the sensu API
const StashesURI string = "/stashes"

// Stash represents a Stash object in the Sensu API
type Stash struct {
	Path    string                 `json:"path"`
	Content map[string]interface{} `json:"content"`
	Expire  string                 `json:"expire"`
}

// GetStashes gets all stashes
func (c *API) GetStashes(out interface{}) (*http.Response, error) {
	resp, err := c.get(StashesURI, out)
	return resp, err
}

// GetStash gets a stash by path
func (c *API) GetStash(out interface{}, path string) (*http.Response, error) {
	s := []string{StashesURI, path}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// RemoveStash removes a stash by the given path
func (c *API) RemoveStash(path string) (*http.Response, error) {
	s := []string{StashesURI, path}
	uri := strings.Join(s, "/")
	resp, err := c.delete(uri)
	return resp, err
}

// CreateStash creates a stash
// POST a Stash object to the API
func (c *API) CreateStash(payload interface{}) (*http.Response, error) {
	resp, err := c.post(StashesURI, payload)
	return resp, err
}
