/*
List and resolve current events. Every event occurrence has a unique ID (random UUID).
https://sensuapp.org/docs/0.20/api-events
*/

package sensu

import (
	"net/http"
	"strings"
)

// EventsURI constant
const EventsURI string = "/events"

// Event represents an Event object in the Sensu API
type Event struct {
	ID          string `json:"id"`
	Client      Client `json:"client"`
	Check       Check  `json:"check"`
	Occurrences int    `json:"occurrences"`
	Action      string `json:"action"`
}

// Resolution represents a Resolution object in the Sensu API
type Resolution struct {
	ClientName string `json:"client"`
	CheckName  string `json:"check"`
}

// GetEvents gets all events
func (c *API) GetEvents(out interface{}) (*http.Response, error) {
	resp, err := c.get(EventsURI, out)
	return resp, err
}

// GetEventsByClient gets all events for a particular client name
func (c *API) GetEventsByClient(out interface{}, clientName string) (*http.Response, error) {
	s := []string{EventsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// GetEventByClientCheck gets an event by client name and check name
func (c *API) GetEventByClientCheck(out interface{}, clientName string, checkName string) (*http.Response, error) {
	s := []string{EventsURI, clientName, checkName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}

// ResolveEvent resolves an event given the client name and check name
func (c *API) ResolveEvent(clientName, checkName string) (*http.Response, error) {
	s := []string{EventsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.delete(uri)
	return resp, err
}

// ResolveEventPost resolves an event given the client name and check name in a POST payload
func (c *API) ResolveEventPost(payload interface{}) (*http.Response, error) {
	resp, err := c.post(EventsURI, payload)
	return resp, err
}
