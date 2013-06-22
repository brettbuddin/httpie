package httpie

import (
    "net/http"
)

func NewClient(authorizer Authorizer) *Client {
    return &Client{
        authorizer: authorizer,
    }
}

type Client struct {
    endpoint   Endpoint
    authorizer Authorizer
}

func (c *Client) Request(end Endpoint) (*http.Response, error) {
    client := &http.Client{}
    req    := &http.Request{Header: http.Header{}}

    end.ApplyTo(req)

    if c.authorizer != nil {
        c.authorizer.Authorize(req)
    }

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}
