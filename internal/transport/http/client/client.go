// Package client sending requests to network services
package client

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

// Client for client
type Client interface {
	Post(ctx context.Context, reqURL string, headers map[string]string, in, out interface{}) error
	Get(ctx context.Context, reqURL string, headers map[string]string, in, out interface{}) error
}

func New(hostname string, httpClient *http.Client, opts ...Option) Client {
	clnt := http.Client(*httpClient)
	c := &client{}
	c.hostname = hostname
	c.httpClient = &clnt

	defaultOptions := options{}
	for _, o := range opts {
		o.apply(&defaultOptions)
	}

	c.external = defaultOptions.external

	return c
}

type client struct {
	hostname   string
	httpClient *http.Client
	external   bool
}

func (c *client) Get(ctx context.Context, reqURL string, headers map[string]string, in, out interface{}) error {
	var err error

	reqURL = c.requestURL(reqURL)
	log.Println("Sending GET request to", reqURL)
	requestBody, err := c.prepareRequestBody(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, strings.NewReader(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	return c.parseResponseBody(responseBody, httpResponse.StatusCode, out)
}

func (c *client) Post(ctx context.Context, reqURL string, headers map[string]string, in, out interface{}) error {
	var err error

	reqURL = c.requestURL(reqURL)
	log.Println("Sending POST request to", reqURL)

	requestBody, err := c.prepareRequestBody(in)
	if err != nil {
		log.Println("Error marshalling request data to JSON")
		return err
	}
	log.Println("Request Body", requestBody)
	req, err := http.NewRequestWithContext(ctx, "POST", reqURL, strings.NewReader(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println("Error reading response body")
		return err
	}
	log.Println("Response: ", string(responseBody))
	return c.parseResponseBody(responseBody, httpResponse.StatusCode, out)
}

func (c *client) prepareRequestBody(req interface{}) (string, error) {
	requestByteJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	bodyStr := string(requestByteJSON)
	if bodyStr == "null" {
		bodyStr = ""
	}

	return bodyStr, nil
}

func (c *client) parseResponseBody(body []byte, statusCode int, out interface{}) error {
	if c.external {
		if statusCode < 200 || statusCode >= 300 {
			return &ErrInvalidResponse{statusCode, statusCode, string(body)}
		}
		if out == nil {
			return nil
		}
		return json.Unmarshal(body, out)
	}

	resp := &response{}
	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return &ErrInvalidResponse{statusCode, resp.Error.Code, resp.Error.Message}
	}

	if out == nil {
		return nil
	}

	if resp.Data == nil {
		return json.Unmarshal(body, out)
	}

	outJSON, err := json.Marshal(resp.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(outJSON, out); err != nil {
		return err
	}

	return nil
}

func (c *client) requestURL(apiMethod string) string {
	params := []string{c.hostname, apiMethod}

	return strings.Join(params, "/")
}
