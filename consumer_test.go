package httpie

import (
    "testing"
    "strings"
    "bufio"
    "bytes"
)

var delimData = []string{
    "Hello",
    "World",
    "This",
    "Is",
    "A",
    "Test",
};

func TestDelim(t *testing.T) {
    data  := []byte(strings.Join(delimData, "\n") + "\n")
    delim := Delimeter{'\n'}

    reader := bufio.NewReader(bytes.NewBuffer(data))

    out := []string{}
    for _, seg := range delimData {
        b, _ := delim.Consume(reader)

        if string(b) == seg {
            out = append(out, string(b))
        }
    }

    if len(out) != len(delimData) {
        t.Errorf("Delimeter consumer output doesn't match: actual=%s expected=%s", out, delimData)
    }
}
