package gst

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURLv1 = "https://api.xxx.com/v1"

func newTransport() *http.Client {
	return &http.Client{}
}

// Client is XXX API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	AuthTToken string
}

// NewClient creates new XXX API client.
func NewClient(authToken string) *Client {
	return &Client{
		BaseURL:    apiURLv1,
		HTTPClient: newTransport(),
		AuthTToken: authToken,
	}
}

type ErrorResponse struct {
	Error *struct {
		Code    *int   `json:"code,omitempty"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AuthTToken))

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			return fmt.Errorf("error, status code: %d", res.StatusCode)
		}
		return fmt.Errorf("error, status code: %d, message: %s", res.StatusCode, errRes.Error.Message)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, suffix)
}
