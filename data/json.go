package dataimport

import (
	"encoding/json"
	"io"
)

// FromJSON decodes a input from JSON to given interface
func FromJSON(i interface{}, reader io.Reader) error {
	d := json.NewDecoder(reader)
	return d.Decode(i)
}

// ToJSON encodes a input from JSON to given interface
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
