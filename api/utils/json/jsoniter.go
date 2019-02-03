package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var (
	json jsoniter.API
)

// Marshal 字面意思
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal 字面意思
func Unmarshal(b []byte, v interface{}) error {

	return json.Unmarshal(b, v)
}

// MarshalToString 字面意思
func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

// UnmarshalFromString 字面意思
func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

// NewDecoder 字面意思
func NewDecoder(reader io.ReadCloser) *jsoniter.Decoder {
	return json.NewDecoder(reader)
}

func init() {
	json = jsoniter.ConfigFastest
}
