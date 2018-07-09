package json

import (
	"io"
	"log"
)

var (
	j Base
)

// Decoder 解析器接口定义
type Decoder interface {
	Decode(v interface{}) error
}

// Base json 模块 接口定义
type Base interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	NewDecoder(reader io.ReadCloser) Decoder
}

//InitJSON 创建并注册json解释器
func InitJSON(b Base) {
	if b == nil {
		log.Panic("must set a base json!")
	}
	j = b
}

// Marshal json编码
func Marshal(v interface{}) (string, error) {
	data, err := j.Marshal(v)
	return string(data), err
}

// Unmarshal json解码
func Unmarshal(data string, v interface{}) error {
	return j.Unmarshal([]byte(data), v)
}

// NewDecoder 生成解析器
func NewDecoder(reader io.ReadCloser) Decoder {
	return j.NewDecoder(reader)
}
