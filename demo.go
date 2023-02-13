package gst

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type (
	RequestData  struct{}
	ResponseData struct{}
)

// Get Demo
func (c *Client) GetDemo(ctx context.Context) (resp ResponseData, err error) {
	req, err := http.NewRequest("GET", c.fullURL("/get_demo"), nil)
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &resp)
	return
}

// Post Demo
func (c *Client) PostDemo(ctx context.Context, request RequestData) (resp ResponseData, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", c.fullURL("/post_demo"), bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	err = c.sendRequest(req, &resp)
	return
}
