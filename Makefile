PACKAGES=$(shell go list ./...)
PACKAGES_INTEGRATION=$(shell go list ./... | grep  integration_test)
PACKAGES_UNITTEST=$(shell go list ./... | grep -v integration_test)
export GO111MODULE = on

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./lite/*/statik.go" -not -path "*.pb.go" | xargs goimports -w -local github.com/bianjieai/iritamod-sdk-go

test-integration:
	cd integration_test/scripts/ && sh build.sh && sh start.sh
	sleep 5
	@go test -v $(PACKAGES_INTEGRATION)
	cd integration_test/scripts/ && sh clean.sh
