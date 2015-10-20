/*
Author: Jeff LaPlante - jeff@jefflaplante.com
License: GPLv3

Sensu API Library for Golang
Supports Sensu 0.20.0. Other versions may have unknown breaking changes.
*/

package sensu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Config is used to configure the creation of a client
type Config struct {
	Address    string
	Scheme     string
	HTTPClient *http.Client
}

// API Client is used as a handle for all client methods
type API struct {
	config Config
}

// DefaultConfig sets up a default configuration struct
func DefaultConfig() *Config {
	config := &Config{
		Scheme:     "http",
		Address:    "127.0.0.1:4567",
		HTTPClient: http.DefaultClient,
	}
	return config
}

// NewAPIClient gets a new Sensu API client
func NewAPIClient(config *Config) (*API, error) {
	defConfig := DefaultConfig()

	if len(config.Scheme) == 0 {
		config.Scheme = defConfig.Scheme
	}

	if len(config.Address) == 0 {
		config.Address = defConfig.Address
	}

	if config.HTTPClient == nil {
		config.HTTPClient = defConfig.HTTPClient
	}

	apiClient := &API{
		config: *config,
	}

	return apiClient, nil
}

// Build a http request
func (c *API) buildRequest(method, path string) (*http.Request, error) {
	url := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.Address,
		Path:   path,
	}

	req, err := http.NewRequest(method, url.String(), nil)
	return req, err
}

// Build a POST http request
func (c *API) buildRequestPOST(path string, payload interface{}) (*http.Request, error) {
	method := "POST"
	url := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.Address,
		Path:   path,
	}

	// Encode payload struct into JSON and create a reader for it
	encodedPayload, err := json.Marshal(payload)
	payloadReader := bytes.NewReader(encodedPayload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), payloadReader)
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

// Send the request to the server
func (c *API) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.config.HTTPClient.Do(req)
	return resp, err
}

// Decode JSON payload
func jsonDecode(out interface{}, data io.ReadCloser) error {
	d := json.NewDecoder(data)
	error := d.Decode(out)
	return error
}

// Generic GET request. Decoded JSON is set in the out interface{} passed in.
func (c *API) get(uri string, out interface{}) (*http.Response, error) {
	request, _ := c.buildRequest("GET", uri)
	resp, err := c.doRequest(request)

	if err != nil {
		return nil, err
	}

	if out != nil {
		err = jsonDecode(out, resp.Body)
	}

	return resp, err
}

// Generic POST request.
func (c *API) post(uri string, payload interface{}) (*http.Response, error) {
	request, _ := c.buildRequestPOST(uri, payload)
	resp, err := c.doRequest(request)

	if err != nil {
		return nil, err
	}
	return resp, err
}

// Generic DELETE request.
func (c *API) delete(uri string) (*http.Response, error) {
	request, _ := c.buildRequest("DELETE", uri)
	resp, err := c.doRequest(request)

	if err != nil {
		return nil, err
	}
	return resp, err
}
