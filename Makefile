GOTEST_PACKAGES = $(shell go list ./... | egrep -v '(pkg|cmd)')

gobuild:
	go build -o ./bin/tz cmd/tz/main.go

gotest: gobuild
	go test -race -v $(GOTEST_PACKAGES)

golint:
	golangci-lint run -v