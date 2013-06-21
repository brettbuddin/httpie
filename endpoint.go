package httpie

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "bytes"
)

type Endpoint interface {
    ApplyTo(*http.Request)
}

type Get struct {
    *url.URL
}

func (g Get) ApplyTo(req *http.Request) {
    req.Method = "GET"
    req.URL    = g.URL
}

type Post struct {
    *url.URL
    Body []byte
    ContentType string
}

func (p Post) ApplyTo(req *http.Request) {
    req.Method        = "POST"
    req.URL           = p.URL
    req.Body          = ioutil.NopCloser(bytes.NewBuffer(p.Body))
    req.ContentLength = int64(len(p.Body))
    req.Header.Set("Content-Type", p.ContentType)
}

type Put struct {
    *url.URL
    Body []byte
    ContentType string
}

func (p Put) ApplyTo(req *http.Request) {
    req.Method        = "PUT"
    req.URL           = p.URL
    req.Body          = ioutil.NopCloser(bytes.NewBuffer(p.Body))
    req.ContentLength = int64(len(p.Body))
    req.Header.Set("Content-Type", p.ContentType)
}

type Delete struct {
    *url.URL
}

func (d Delete) ApplyTo(req *http.Request) {
    req.Method = "DELETE"
    req.URL    = d.URL
}
