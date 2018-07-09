package json

import (
	"io"

	"github.com/json-iterator/go"
)

// JSONiter jsoniter base 实现
type JSONiter struct {
	json jsoniter.API
}

// Marshal 字面意思
func (j *JSONiter) Marshal(v interface{}) ([]byte, error) {
	return j.json.Marshal(v)
}

// Unmarshal 字面意思
func (j *JSONiter) Unmarshal(b []byte, v interface{}) error {
	return j.json.Unmarshal(b, v)
}

// NewDecoder 字面意思
func (j *JSONiter) NewDecoder(reader io.ReadCloser) Decoder {
	return j.json.NewDecoder(reader)
}

// NewJSONiter 创建JSONiter 实例
func NewJSONiter() *JSONiter {
	return &JSONiter{
		json: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}
