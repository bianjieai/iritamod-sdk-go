module github.com/bianjieai/iritamod-sdk-go

go 1.16

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

require (
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/core-sdk-go v0.0.0-20211101034133-8664fa1c1205
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/prometheus/common v0.23.0 // indirect
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.12
	github.com/tendermint/tm-db v0.6.4 // indirect
	google.golang.org/genproto v0.0.0-20210916144049-3192f974c780
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.3.0
)
