module github.com/bianjieai/iritamod-sdk-go/perm

go 1.16

require (
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/irisnet/core-sdk-go v0.0.0-20210817104504-bd2c112847e9
	google.golang.org/genproto v0.0.0-20210916144049-3192f974c780 // indirect
	google.golang.org/grpc v1.40.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/irisnet/core-sdk-go => D:\desktop\tibc\core-sdk-go
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
