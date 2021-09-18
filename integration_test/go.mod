module github.com/bianjieai/iritamod-sdk-go/integration

go 1.16

replace (
	github.com/bianjieai/iritamod-sdk-go/identity => D:\desktop\tmp\iritamod-sdk-go\identity
	github.com/bianjieai/iritamod-sdk-go/nft => D:\desktop\tmp\iritamod-sdk-go\nft
	github.com/bianjieai/iritamod-sdk-go/node => D:\desktop\tmp\iritamod-sdk-go\node
	github.com/bianjieai/iritamod-sdk-go/oracle => D:\desktop\tmp\iritamod-sdk-go\oracle
	github.com/bianjieai/iritamod-sdk-go/params => D:\desktop\tmp\iritamod-sdk-go\params
	github.com/bianjieai/iritamod-sdk-go/perm => D:\desktop\tmp\iritamod-sdk-go\perm
	github.com/bianjieai/iritamod-sdk-go/record => D:\desktop\tmp\iritamod-sdk-go\record
	github.com/bianjieai/iritamod-sdk-go/service => D:\desktop\tmp\iritamod-sdk-go\service
	github.com/bianjieai/iritamod-sdk-go/token => D:\desktop\tmp\iritamod-sdk-go\token
	github.com/bianjieai/iritamod-sdk-go/wasm => D:\desktop\tmp\iritamod-sdk-go\wasm
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/core-sdk-go => D:\desktop\tibc\core-sdk-go
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

require (
	github.com/irisnet/core-sdk-go v0.0.0-20210817104504-bd2c112847e9
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
)
