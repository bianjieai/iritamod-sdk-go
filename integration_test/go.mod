module github.com/bianjieai/iritamod-sdk-go/integration

go 1.16

require (
	github.com/bianjieai/iritamod-sdk-go/identity v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/nft v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/node v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/oracle v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/params v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/perm v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/record v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/service v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/token v0.0.0-20210922060017-08a51dc9e679
	github.com/bianjieai/iritamod-sdk-go/wasm v0.0.0-20210922060017-08a51dc9e679
	github.com/irisnet/core-sdk-go v0.0.0-20210922011537-f1a5093df21b
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
