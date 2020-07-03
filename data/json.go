package dataimport

import (
	"encoding/json"
	"io"
)

func FromJson(i interface{}, reader io.Reader) error  {
	d := json.NewDecoder(reader)
	return d.Decode(i)
}

func ToJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
