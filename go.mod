module github.com/itsfunny/go-cell

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.10.8
	github.com/okex/exchain-ethereum-compatible v1.1.1-0.20220106042715-f20163fbb4af
)

require (
	github.com/ChengjinWu/gojson v0.0.0-20181113073026-04749cc2d015
	github.com/emirpasic/gods v1.12.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/go-openapi/spec v0.20.4
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.1.5
	github.com/gorilla/websocket v1.4.2
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d
	github.com/sasha-s/go-deadlock v0.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/swag v1.7.8
	github.com/tendermint/tendermint v0.33.9
	github.com/tidwall/gjson v1.12.1
	go.uber.org/fx v1.16.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	google.golang.org/grpc v1.28.1
)

replace (
	github.com/buger/jsonparser => github.com/buger/jsonparser v1.0.0 // imported by nacos-go-sdk, upgraded to v1.0.0 in case of a known vulnerable bug
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-oec3
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-exchain6
)
