package httpie

import (
    "bufio"
    "bytes"
)

var (
    NewLine        = Delimeter{'\n'}
    CarriageReturn = Delimeter{'\r'}
    Space          = Delimeter{' '}
    Comma          = Delimeter{','}
)

type Consumer interface {
    Consume(*bufio.Reader) ([]byte, error)
}

type Delimeter struct {
    delim byte
}

func (d Delimeter) Consume(reader *bufio.Reader) ([]byte, error) {
    b, err := reader.ReadBytes(d.delim)

    if err != nil {
        return nil, err
    }

    return bytes.TrimSpace(b), nil
}
