package types

import (
	"fmt"
	"github.com/fxamacker/cbor/v2"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"testing"
)

func TestCbor(t *testing.T) {
	data := types2.Envelope{}
	data.Header = &types2.EnvelopeHeader{
		Flag:       1,
		Length:     2,
		Protocol:   "asd",
		SequenceId: "vv",
	}
	ret, er := cbor.Marshal(data)
	fmt.Println(er)
	fmt.Println(string(ret))
	var unasd types2.Envelope
	er = cbor.Unmarshal(ret, &unasd)
	fmt.Println(er)
	fmt.Println(unasd.Header.Protocol)
}
