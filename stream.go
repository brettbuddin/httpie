package httpie

import (
    "net/http"
    "bufio"
    "fmt"
    "errors"
)

// NewStream returns a Stream
func NewStream(endpoint Endpoint, auth Authorizer, consumer Consumer) *Stream {
    return &Stream{
        endpoint:   endpoint,
        authorizer: auth,
        consumer:   consumer,
        data:       make(chan []byte, 50),
    }
}

type Stream struct {
    data        chan []byte
    endpoint    Endpoint
    authorizer  Authorizer
    consumer    Consumer
}

// Connect starts the stream
func (s *Stream) Connect() {
    resp, err := s.connect()
    if err != nil {
        return
    }

    s.consume(resp)
}

// Data returns a channel that chunks of the
// feed will be communicated upon
func (s *Stream) Data() (chan []byte) {
    return s.data
}

func (s *Stream) connect() (*http.Response, error) {
    client := &http.Client{}
    req    := &http.Request{Header: http.Header{}}

    s.endpoint.ApplyTo(req)
    if s.authorizer != nil {
        s.authorizer.Authorize(req)
    }

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("Status code received: %s", resp.StatusCode))
    }

    return resp, nil
}

func (s *Stream) consume(resp *http.Response) {
    reader := bufio.NewReader(resp.Body)

    var (
        b []byte
        err error
    )

    for {
        b, err = s.consumer.Consume(reader)

        if err != nil {
            resp.Body.Close()

            if resp, err = s.connect(); err != nil {
                continue
            }

            reader = bufio.NewReader(resp.Body)
        }

        s.data <- b
    }
}
