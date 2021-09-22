module github.com/bianjieai/iritamod-sdk-go/oracle

go 1.16

require (
	github.com/bianjieai/iritamod-sdk-go/service v0.0.0-20210922060017-08a51dc9e679
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/core-sdk-go v0.0.0-20210922011537-f1a5093df21b
	google.golang.org/genproto v0.0.0-20210916144049-3192f974c780
	google.golang.org/grpc v1.40.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113

)
