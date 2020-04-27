// Package json replaces encoding/json with jsoniter.
package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Decoder = jsoniter.Decoder

func NewDecoder(r io.Reader) *Decoder { return json.NewDecoder(r) }

type Encoder = jsoniter.Encoder

func NewEncoder(w io.Writer) *Encoder {
	return json.NewEncoder(w)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
