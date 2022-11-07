package codec

import (
	"github.com/itsfunny/go-cell/base/common/enums"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/base"
	"github.com/itsfunny/go-cell/component/codec/impl"
	"github.com/itsfunny/go-cell/component/codec/types"
	"go.uber.org/fx"
)

var (
	CodecModule = fx.Options(
		fx.Provide(NewCodecComponent),
	)
)

type CodecComponent struct {
	*base.BaseComponent

	cdc types.Codec
}

func NewCodecComponent() *CodecComponent {
	ret := &CodecComponent{
		cdc: impl.NewDefaultCodec(),
	}
	ret.BaseComponent = base.NewBaseComponent(enums.CodecModule, ret)

	return ret
}

func (c *CodecComponent) OnStart(ctx *services.StartCTX) error {
	return nil
}

func (c *CodecComponent) GetCodec() types.Codec {
	return c.cdc
}

func (c *CodecComponent) Marshal(data interface{}) ([]byte, error) {
	return c.cdc.Marshal(data)
}
func (c *CodecComponent) UnMarshal(data []byte, ret interface{}) error {
	return c.cdc.Unmarshal(data, ret)
}
