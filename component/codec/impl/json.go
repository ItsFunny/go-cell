package impl

import (
	"github.com/fxamacker/cbor/v2"
	_ "github.com/fxamacker/cbor/v2"
	"github.com/itsfunny/go-cell/component/codec/types"
)

var (
	_ types.Codec = (*DefaultCodec)(nil)
)

type DefaultCodec struct {
}

func NewDefaultCodec() *DefaultCodec {
	return &DefaultCodec{}
}

func (d *DefaultCodec) Marshal(data interface{}) ([]byte, error) {
	if m, ok := data.(types.Marshaller); ok {
		return m.Marshal()
	}
	return cbor.Marshal(data)
}

func (d *DefaultCodec) Unmarshal(data []byte, ret interface{}) error {
	if un, ok := ret.(types.Unmarshaler); ok {
		return un.Unmarshal(data)
	}
	return cbor.Unmarshal(data, ret)
}
