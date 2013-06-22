package httpie

import (
    "net/http"
)

type Authorizer interface {
    Authorize(*http.Request)
}

type BasicAuth struct {
    Username, Password string
}

func (b BasicAuth) Authorize(req *http.Request) {
    req.SetBasicAuth(b.Username, b.Password)
}

type HeaderAuth struct {
    Auth string
}

func (h HeaderAuth) Authorize(req *http.Request) {
    req.Header.Set("Authorization", h.Auth)
}
