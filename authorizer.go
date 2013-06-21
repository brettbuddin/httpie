package httpie

import (
    "net/http"
)

type Authorizer interface {
    Authorize(*http.Request)
}

type BasicAuth struct {
    username, password string
}

func (b BasicAuth) Authorize(req *http.Request) {
    req.SetBasicAuth(b.username, b.password)
}

type HeaderAuth struct {
    auth string
}

func (h HeaderAuth) Authorize(req *http.Request) {
    req.Header.Set("Authorization", h.auth)
}
