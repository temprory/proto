package proto

import (
	"github.com/json-iterator/go"
	"github.com/temprory/net"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

var (
	DefaultCodec IRpcCodec = &JsonCodec{}
)

func SetCodec(c IRpcCodec) {
	DefaultCodec = c
}

type IRpcCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type JsonCodec struct{}

func (c *JsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return DefaultCodec.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return DefaultCodec.Unmarshal(data, v)
}

func NewMessage(cmd uint32, v interface{}) *net.Message {
	data, _ := Marshal(v)
	return net.NewMessage(cmd, data)
}
